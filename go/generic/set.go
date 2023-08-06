package generic

import "sync"

// TODO: Reimplement using sync.Map
type Set[T comparable] struct {
	items map[T]struct{}
	m     sync.RWMutex
}

func NewSet[T comparable](items []T) *Set[T] {
	s := &Set[T]{
		items: make(map[T]struct{}),
	}
	s.Store(items...)
	return s
}

func NewSetFn[T comparable, K any](items []K, fn func(K) T) *Set[T] {
	s := &Set[T]{
		items: make(map[T]struct{}),
	}
	for _, item := range items {
		s.Store(fn(item))
	}
	return s
}

func (s *Set[T]) Store(items ...T) {
	s.m.Lock()
	defer s.m.Unlock()
	for _, item := range items {
		s.items[item] = struct{}{}
	}
}

func (s *Set[T]) Has(item T) bool {
	s.m.RLock()
	defer s.m.RUnlock()
	_, ok := s.items[item]
	return ok
}

func (s *Set[T]) Delete(item T) {
	s.m.Lock()
	defer s.m.Unlock()
	delete(s.items, item)
}

func (s *Set[T]) Len() int {
	s.m.RLock()
	defer s.m.RUnlock()
	return len(s.items)
}
