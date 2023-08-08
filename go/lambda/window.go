package lambda

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
	precedingIdx := idx - preceding
	if precedingIdx < 0 {
		precedingIdx = 0
	}
	followingIdx := idx + following + 1
	if followingIdx > len(items) {
		followingIdx = len(items)
	}
	return items[precedingIdx:followingIdx:followingIdx]
}

type Groupable interface {
	Position()
}
