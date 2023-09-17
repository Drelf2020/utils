package utils

// If expr is true, return a.
//
// Otherwise return b.
func Ternary[T any](expr bool, a, b T) T {
	if expr {
		return a
	}
	return b
}

// If expr is true, set a to the value of b.
//
// Otherwise do nothing.
func Update[T any](expr bool, a *T, b T) {
	if expr {
		*a = b
	}
}

// Slice Filter
func Filter[S ~[]E, E any](s S, f func(E) bool) S {
	r := make(S, 0, len(s))
	for _, e := range s {
		if f(e) {
			r = append(r, e)
		}
	}
	return r
}

// Type must be a pointer, channel, func, interface, map, or slice type.
//
// But we don't consider map and interface.
type CanNil[T any] interface {
	~*T | ~chan T | ~func() | ~func(T) | ~func() T | ~[]T
}

// Return a copy of s without nil elements.
//
// What you need is to type T which can't be infer...
//
// If you still want auto infer T's type,
// you can use the functions below,
// like NotNilPtr / NotNilChan / NotNilSlice
func NotNil[T any, E CanNil[T], S ~[]E](s S) S {
	return Filter(s, func(e E) bool { return e != nil })
}

func NotNilPtr[T any, E ~*T, S ~[]E](s S) S {
	return NotNil[T](s)
}

func NotNilChan[T any, E ~chan T, S ~[]E](s S) S {
	return NotNil[T](s)
}

func NotNilSlice[T any, E ~[]T, S ~[]E](s S) S {
	return NotNil[T](s)
}

// Like the map() function in Python.
func Map[T, V any](f func(T) V, iter []T) []V {
	v := make([]V, 0, len(iter))
	for _, i := range iter {
		v = append(v, f(i))
	}
	return v
}

// Like the reduce() function in Python.
func Reduce[T, V any](f func(T, V) T, v []V, s T) T {
	for _, e := range v {
		s = f(s, e)
	}
	return s
}
