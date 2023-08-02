package concurrency

import (
	"context"

	"golang.org/x/sync/errgroup"
)

func Do[T any](concurrencyLimit int, fn func(T) error, params ...T) error {
	wg := errgroup.Group{}
	wg.SetLimit(concurrencyLimit)
	for i := range params {
		wg.Go(doRequest(fn, params[i]))
	}
	return wg.Wait()
}

func doRequest[T any](fn func(T) error, param T) func() error {
	return func() error {
		return fn(param)
	}
}

func DoContext[T any](ctx context.Context, concurrencyLimit int, fn func(context.Context, T) error, params ...T) error {
	wg, cctx := errgroup.WithContext(ctx)
	wg.SetLimit(concurrencyLimit)
	for i := range params {
		wg.Go(doRequestContext(cctx, fn, params[i]))
	}
	return wg.Wait()
}

func doRequestContext[T any](ctx context.Context, fn func(context.Context, T) error, param T) func() error {
	return func() error {
		errChan := make(chan error)
		go func() {
			errChan <- fn(ctx, param)
		}()
		select {
		case <-ctx.Done():
			return ctx.Err()
		case err := <-errChan:
			return err
		}
	}
}

type ResponseHandler[T, V any] interface {
	Value() (id T, result V, err error)
}

type result[T, V any] struct {
	err error
	val V
	id  T
}

func newResult[T, V any](id T, resp V, err error) ResponseHandler[T, V] {
	return &result[T, V]{err: err, id: id, val: resp}
}

func (r *result[T, V]) Value() (id T, result V, err error) {
	return r.id, r.val, r.err
}

func DoChan[T, V any](limit, sz int, fn func(T) (V, error), params ...T) chan ResponseHandler[T, V] {
	responseChan := make(chan ResponseHandler[T, V], sz)
	go func() {
		defer close(responseChan)
		wg := errgroup.Group{}
		wg.SetLimit(limit)
		for i := range params {
			wg.Go(doRequestChan(responseChan, fn, params[i]))
		}
		wg.Wait()
	}()
	return responseChan
}

func doRequestChan[T, V any](ch chan ResponseHandler[T, V], fn func(T) (V, error), param T) func() error {
	return func() error {
		res, err := fn(param)
		ch <- newResult(param, res, err)
		return err
	}
}
