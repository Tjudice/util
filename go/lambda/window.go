package lambda

import "github.com/tjudice/util/go/generic"

func Window[T any](preceding, following int, items []T) [][]T {
	windows := make([][]T, 0, len(items))
	for i := range items {
		windows = append(windows, makeWindow(i, preceding, following, items))
	}
	return windows
}

func WindowFn[T any](preceding, following int, items []T, fn func([]T)) {
	for i := range items {
		fn(makeWindow(i, preceding, following, items))
	}
}

func makeWindow[T any](idx, preceding, following int, items []T) []T {
	return items[generic.Max(idx-preceding, 0):generic.Min(idx+following+1, len(items))]
}

type Groupable interface {
	Position()
}
