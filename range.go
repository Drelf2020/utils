package utils

// It is just a typed Channel
type R[T int | int64 | float64] chan T

func (r R[T]) List() (ls []T) {
	for v := range r {
		ls = append(ls, v)
	}
	return
}

// A Range(start, stop[, step]) function return slice like python
func Range[T int | int64 | float64](stop T, options ...T) R[T] {
	var start, step, alpha T = 0, 1, 1
	switch len(options) {
	case 2:
		step = options[1]
		fallthrough
	case 1:
		start = stop
		stop = options[0]
	}
	if step < 0 {
		alpha = -1
	}

	ch := make(R[T])
	go func(ch R[T]) {
		defer func() { recover() }()
		for i := start; (i-stop)*alpha < 0; i += step {
			ch <- i
		}
		close(ch)
	}(ch)
	return ch
}
