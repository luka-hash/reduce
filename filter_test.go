// This file contains code that is a combination of two sources:
// 1. Original code under the MIT Licence:
//    Copyright © Luka Ivanović
// 2. Code derived from the [filter](https://github.com/robpike/filter) Project (BSD-style licensed code):
//    Copyright © The Go Authors
// See LICENCE for details.

package reduce

import (
	"reflect"
	"testing"
)

func isEven(a int) bool {
	return a%2 == 0
}

func isOdd(a int) bool {
	return a%2 == 1
}

func isEvenString(a string) bool {
	return a[0]%2 == 0
}

func TestChoose(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	expect := []int{2, 4, 6, 8}
	result := Choose(a, isEven)
	if !reflect.DeepEqual(expect, result) {
		t.Fatalf("Choose failed: expect %v got %v", expect, result)
	}
}

func TestDrop(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expect := []int{2, 4, 6, 8, 10}
	result := Drop(a, isOdd)
	if !reflect.DeepEqual(expect, result) {
		t.Fatalf("Drop failed: expect %v got %v", expect, result)
	}
}

func TestChooseInPlaceString(t *testing.T) {
	a := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	expect := []string{"2", "4", "6", "8"}
	ChooseInPlace(&a, isEvenString)
	if !reflect.DeepEqual(expect, a) {
		t.Fatalf("ChooseInPlace failed: expect %v got %v", expect, a)
	}
}

func TestDropInPlaceString(t *testing.T) {
	a := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	expect := []string{"1", "3", "5", "7", "9"}
	DropInPlace(&a, isEvenString)
	if !reflect.DeepEqual(expect, a) {
		t.Fatalf("DropInPlace failed: expect %v got %v", expect, a)
	}
}
