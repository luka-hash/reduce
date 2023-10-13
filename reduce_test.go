// This file contains code that is a combination of two sources:
// 1. Original code under the MIT Licence:
//    Copyright © Luka Ivanović
// 2. Code derived from the [filter](https://github.com/robpike/filter) Project (BSD-style licensed code):
//    Copyright © The Go Authors
// See LICENCE for details.

package reduce

import (
	"testing"
)

func mul(a, b int) int {
	return a * b
}

func fac(n int) int {
	if n <= 1 {
		return 1
	}
	return n * fac(n-1)
}

func TestReduce(t *testing.T) {
	const size = 20
	a := make([]int, size)
	for i := range a {
		a[i] = i + 1
	}
	for i := 1; i < size; i++ {
		out := Reduce(a[:i], mul, 1)
		expect := fac(i)
		if expect != out {
			t.Fatalf("expected %d got %d", expect, out)
		}
	}
}
