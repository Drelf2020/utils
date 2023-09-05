package utils

// 过滤数组
func Filter[S ~[]E, E any](s S, f func(E) bool) S {
	r := make(S, 0, len(s))
	for _, e := range s {
		if f(e) {
			r = append(r, e)
		}
	}
	return r
}

// 类似 Python 的 map 函数
func Map[T, V any](f func(T) V, iter []T) []V {
	v := make([]V, 0, len(iter))
	for _, i := range iter {
		v = append(v, f(i))
	}
	return v
}

// 类似 Python 的 reduce 函数
func Reduce[T, V any](f func(T, V) T, v []V, s T) T {
	for _, e := range v {
		s = f(s, e)
	}
	return s
}
