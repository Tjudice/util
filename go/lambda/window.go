package lambda

// type WindowItem interface {
// 	constraints.Integer | constraints.Float
// }

// // Assumes input array is sorted - Not sure its neccesary to add another function param?
// func WindowSlice[T any, Unit WindowItem](start, end, length Unit, items []T, inside func(item T, i, j Unit)) [][]T {
// 	return nil
// }

func Window[T any](preceding, following int, items []T) [][]T {
	windows := make([][]T, 0, len(items))
	for i := range items {
		precedingIdx := i - preceding
		if precedingIdx < 0 {
			precedingIdx = 0
		}
		followingIdx := i + following + 1
		if followingIdx > len(items) {
			followingIdx = len(items)
		}
		windows = append(windows, items[precedingIdx:followingIdx])
	}
	return windows
}
