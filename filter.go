// Copyright © 2023 Luka Ivanović
// This code is licensed under MIT licence (see LICENCE for details)

package reduce

// This could and should just be:
// 	func Filter[T any](xs []T, test func(T) bool) []T
// 	func FilterInPlace[T any](xs *[]T, test func(T) bool)

func Choose[T any](xs []T, test func(T) bool) []T {
	res := make([]T, len(xs))
	l := 0
	for i := range xs {
		if test(xs[i]) {
			res[l] = xs[i]
			l += 1
		}
	}
	return res[:l]
}

func ChooseInPlace[T any](xs *[]T, test func(T) bool) {
	l := 0
	for i := range *xs {
		if test((*xs)[i]) {
			(*xs)[l] = (*xs)[i]
			l += 1
		}
	}
	*xs = (*xs)[:l]
}

func Drop[T any](xs []T, test func(T) bool) []T {
	res := make([]T, len(xs))
	l := 0
	for i := range xs {
		if !test(xs[i]) {
			res[l] = xs[i]
			l += 1
		}
	}
	return res[:l]
}

func DropInPlace[T any](xs *[]T, test func(T) bool) {
	l := 0
	for i := range *xs {
		if !test((*xs)[i]) {
			(*xs)[l] = (*xs)[i]
			l += 1
		}
	}
	*xs = (*xs)[:l]
}
