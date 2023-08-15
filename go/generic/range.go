package generic

import "golang.org/x/exp/constraints"

type Rangeable interface {
	constraints.Integer
}

type Range[T Rangeable] struct {
	Start T
	End   T
}

// Floating point ranges may have some overlap
func DivideRangeInclusive[T Rangeable](start T, end T, size T) []Range[T] {
	if start > end {
		return nil
	}
	if size == 0 {
		return []Range[T]{
			{Start: start, End: end},
		}
	}
	ranges := make([]Range[T], 0, int((end-start)/size)+1)
	curr := start
	incr := size - 1
	for {
		next := Min(curr+incr, end)
		ranges = append(ranges, Range[T]{
			Start: curr,
			End:   next,
		})
		if next >= end {
			break
		}
		curr = curr + incr + 1
	}
	return ranges
}
