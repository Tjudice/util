package lambda

func MapToSlice[K comparable, V any](m map[K]V) []V {
	res := make([]V, 0, len(m))
	for _, v := range m {
		res = append(res, v)
	}
	return res
}

func TransformSlice[T, K any](vals []T, transformer func(T) K) []K {
	n := make([]K, 0, len(vals))
	for _, v := range vals {
		n = append(n, transformer(v))
	}
	return n
}

func SliceToAny[T any](vals []T) []any {
	return TransformSlice(vals, func(v T) any {
		return v
	})
}
