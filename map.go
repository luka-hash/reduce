// Copyright © 2023 Luka Ivanović
// This code is licensed under MIT licence (see LICENCE for details)

package reduce

func Map[T any, S any](xs []T, fn func(T) S) []S {
	res := make([]S, len(xs))
	for i := range xs {
		res[i] = fn(xs[i])
	}
	return res
}

func MapInPlace[T any](xs []T, fn func(T) T) {
	// TODO: this should prob. be protected with a mutex
	for i := range xs {
		xs[i] = fn(xs[i])
	}
}
