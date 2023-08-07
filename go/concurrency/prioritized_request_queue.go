package concurrency

import (
	"context"
)

type PrioritizedRequestQueue[T, V any] struct {
	pq *PriorityQueue[request[T, V]]

	getter func(context.Context, T) (V, error)

	reqs chan struct{}

	done chan struct{}
}

type request[T, V any] struct {
	ctx   context.Context
	param T
	ch    chan *response[V]
}

type response[V any] struct {
	val V
	err error
}

func NewPrioritizedRequestQueue[T, V any](workers int, less func(i T, j T) bool, getter func(context.Context, T) (V, error)) *PrioritizedRequestQueue[T, V] {
	o := &PrioritizedRequestQueue[T, V]{
		pq:   NewPriorityQueue[request[T, V]](workers*10, makeLessFn[T, V](less)),
		reqs: make(chan struct{}, workers),
		done: make(chan struct{}),
	}
	o.run(workers)
	return o
}

func (p *PrioritizedRequestQueue[T, V]) run(workers int) {
	for i := 0; i < workers; i = i + 1 {
		go p.work()
	}
}

func (p *PrioritizedRequestQueue[T, V]) work() {
	for {
		select {
		case <-p.done:
			return
		case <-p.reqs:
			req, ok := p.pq.TryPop()
			if !ok {
				continue
			}
			res, err := p.getter(req.ctx, req.param)
			req.ch <- &response[V]{val: res, err: err}
		}
	}
}

func makeLessFn[T, V any](less func(i T, j T) bool) func(request[T, V], request[T, V]) bool {
	return func(i request[T, V], j request[T, V]) bool {
		return less(i.param, j.param)
	}
}

func (p *PrioritizedRequestQueue[T, V]) Do(ctx context.Context, params T) (V, error) {
	responseChan := make(chan *response[V])
	p.pq.Push(request[T, V]{ctx: ctx, param: params, ch: responseChan})
	go func() {
		p.reqs <- struct{}{}
	}()
	res := <-responseChan
	return res.val, res.err
}
