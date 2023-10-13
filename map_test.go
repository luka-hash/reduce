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

func triple(a int) int {
	return a * 3
}

func tripleString(a string) string {
	return a + a + a
}

func tripleToFloat(a int) float64 {
	return float64(a * 3)
}

func floatToInt(a float64) int {
	return int(a)
}

func TestMap(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	expect := []int{3, 6, 9, 12, 15, 18, 21, 24, 27}
	result := Map(a, triple)
	if !reflect.DeepEqual(expect, result) {
		t.Fatalf("Map failed: expect %v got %v", expect, result)
	}
}

func TestMapString(t *testing.T) {
	a := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	expect := []string{"111", "222", "333", "444", "555", "666", "777", "888", "999"}
	result := Map(a, tripleString)
	if !reflect.DeepEqual(expect, result) {
		t.Fatalf("Map failed: expect %v got %v", expect, result)
	}
}

func TestMapDifferentTypes(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	expect := []int{3, 6, 9, 12, 15, 18, 21, 24, 27}
	result := Map(Map(a, tripleToFloat), floatToInt)
	if !reflect.DeepEqual(expect, result) {
		t.Fatalf("Map failed: expect %v got %v", expect, result)
	}
}

func TestMapInPlace(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	expect := []int{3, 6, 9, 12, 15, 18, 21, 24, 27}
	MapInPlace(a, triple)
	if !reflect.DeepEqual(expect, a) {
		t.Fatalf("Map failed: expect %v got %v", expect, a)
	}
}

func TestMapInPlaceString(t *testing.T) {
	a := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	expect := []string{"111", "222", "333", "444", "555", "666", "777", "888", "999"}
	MapInPlace(a, tripleString)
	if !reflect.DeepEqual(expect, a) {
		t.Fatalf("Map failed: expect %v got %v", expect, a)
	}
}
