package concurrency

import (
	"context"
)

type PriorityClient[T, V any] interface {
	// Less determines the ordering of fetching when the request queue is clogged,
	// i.e. when the number of available workers is 0
	Less(x, y T) bool
	// Fetch fetches the result of a given request. Requets will be fed to Fetch() based on the ordering as defined by Less(T,T)
	Fetch(context.Context, T) (V, error)
}

type PrioritizedRequestHandler[T, V any] struct {
	less  func(T, T) bool
	fetch func(context.Context, T) (V, error)
}

func (p *PrioritizedRequestHandler[T, V]) Less(x, y T) bool {
	return p.less(x, y)
}

func (p *PrioritizedRequestHandler[T, V]) Fetch(ctx context.Context, params T) (V, error) {
	return p.fetch(ctx, params)
}

type PrioritizedRequestQueue[T, V any] struct {
	pq *PriorityQueue[request[T, V]]

	fetcher PriorityClient[T, V]

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

func NewPrioritizedRequestQueue[T, V any](workers int, less func(i T, j T) bool, fetch func(context.Context, T) (V, error)) *PrioritizedRequestQueue[T, V] {
	f := &PrioritizedRequestHandler[T, V]{less: less, fetch: fetch}
	o := &PrioritizedRequestQueue[T, V]{
		pq:      NewPriorityQueue[request[T, V]](workers*10, makeLessFn[T, V](f.Less)),
		reqs:    make(chan struct{}, workers),
		done:    make(chan struct{}),
		fetcher: f,
	}
	o.run(workers)
	return o
}

func NewPrioritizedRequestQueueInterface[T, V any](workers int, fetcher PriorityClient[T, V]) *PrioritizedRequestQueue[T, V] {
	o := &PrioritizedRequestQueue[T, V]{
		pq:      NewPriorityQueue[request[T, V]](workers*10, makeLessFn[T, V](fetcher.Less)),
		reqs:    make(chan struct{}, workers),
		done:    make(chan struct{}),
		fetcher: fetcher,
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
			res, err := p.fetcher.Fetch(req.ctx, req.param)
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
