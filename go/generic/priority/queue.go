package priority

import "github.com/trevorjudice/util/go/generic/heap"

type Queue[T any] struct {
	items []T
	less  func(T, T) bool
}

func (p *Queue[T]) Len() int {
	return len(p.items)
}

func (p *Queue[T]) Less(i, j int) bool {
	return p.less(p.items[i], p.items[j])
}

func (p *Queue[T]) Swap(i, j int) {
	p.items[i], p.items[j] = p.items[j], p.items[i]
}

func (p *Queue[T]) Push(val T) {
	p.items = append(p.items, val)
}

func (p *Queue[T]) Pop() T {
	old := p.items[len(p.items)-1]
	// how to avoid leak here?
	p.items = p.items[0 : len(p.items)-1]
	return old
}

func NewQueue[T any](cap int, less func(T, T) bool) *Queue[T] {
	queue := &Queue[T]{
		items: make([]T, 0, cap),
		less:  less,
	}
	heap.Init[T](queue)
	return queue
}
