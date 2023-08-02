// Modified Generic container/heap implementation

package heap

import "sort"

type Interface[T any] interface {
	sort.Interface
	Push(x T) // add x as element Len()
	Pop() T   // remove and return element Len() - 1.
}

func Init[T any](h Interface[T]) {
	// heapify
	n := h.Len()
	for i := n/2 - 1; i >= 0; i-- {
		down(h, i, n)
	}
}

func Push[T any](h Interface[T], x T) {
	h.Push(x)
	up(h, h.Len()-1)
}

func Pop[T any](h Interface[T]) T {
	n := h.Len() - 1
	h.Swap(0, n)
	down(h, 0, n)
	return h.Pop()
}

func Remove[T any](h Interface[T], i int) T {
	n := h.Len() - 1
	if n != i {
		h.Swap(i, n)
		if !down(h, i, n) {
			up(h, i)
		}
	}
	return h.Pop()
}

func Fix[T any](h Interface[T], i int) {
	if !down(h, i, h.Len()) {
		up(h, i)
	}
}

func up[T any](h Interface[T], j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !h.Less(j, i) {
			break
		}
		h.Swap(i, j)
		j = i
	}
}

func down[T any](h Interface[T], i0, n int) bool {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && h.Less(j2, j1) {
			j = j2 // = 2*i + 2  // right child
		}
		if !h.Less(j, i) {
			break
		}
		h.Swap(i, j)
		i = j
	}
	return i > i0
}
