package algorithms

import "sort"

type SeriesItem interface {
	Position() int64
}

func LocfObservations[T SeriesItem](start, end, interval int64, items []T, fillFromPrevious func(position int64, last T) T) []T {
	if len(items) == 0 {
		return nil
	}
	sort.SliceStable(items, func(i, j int) bool {
		return items[i].Position() < items[j].Position()
	})
	vals := make([]T, 0, (end-start)/interval+1)
	lastIndex := 0
	for i := start; i <= end; i = i + interval {
		for lastIndex < len(items) && items[lastIndex].Position() <= i {
			lastIndex = lastIndex + 1
		}
		vals = append(vals, fillFromPrevious(i, vals[lastIndex]))
	}
	return vals
}
