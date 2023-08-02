package concurrency

import (
	"sync"

	"github.com/tjudice/util/go/generic/priority"
)

type PriorityQueue[T any] struct {
	m     sync.RWMutex
	inner *priority.PriorityQueue[T]
}

func (p *PriorityQueue[T]) Len() int {
	p.m.RLock()
	defer p.m.RUnlock()
	return p.inner.Len()
}

func (p *PriorityQueue[T]) Cap() int {
	p.m.RLock()
	defer p.m.RUnlock()
	return p.inner.Cap()
}

func (p *PriorityQueue[T]) Push(val T) {
	p.m.Lock()
	defer p.m.Unlock()
	p.inner.Push(val)
}

func (p *PriorityQueue[T]) Pop() T {
	p.m.Lock()
	defer p.m.Unlock()
	return p.inner.Pop()
}

func (p *PriorityQueue[T]) TryPop() (T, bool) {
	p.m.Lock()
	defer p.m.Unlock()
	if p.inner.Len() == 0 {
		return *new(T), false
	}
	return p.inner.Pop(), true
}

func NewPriorityQueue[T any](cap int, less func(T, T) bool) *PriorityQueue[T] {
	queue := priority.NewPriorityQueue[T](cap, less)
	return &PriorityQueue[T]{
		inner: queue,
	}
}
