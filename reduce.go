// Copyright © 2023 Luka Ivanović
// This code is licensed under MIT licence (see LICENCE for details)

package reduce

func Reduce[T any](xs []T, fn func(T, T) T, zero T) T {
	if len(xs) == 0 {
		return zero
	} else if len(xs) == 1 {
		return xs[0]
	} else {
		for i := range xs {
			zero = fn(zero, xs[i])
		}
		return zero
	}
}
