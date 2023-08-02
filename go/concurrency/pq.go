package concurrency

import (
	"sync"

	"github.com/trevorjudice/util/go/generic/priority"
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

func NewPriorityQueue[T any](cap int, less func(T, T) bool) *PriorityQueue[T] {
	queue := priority.NewPriorityQueue[T](cap, less)
	return &PriorityQueue[T]{
		inner: queue,
	}
}
