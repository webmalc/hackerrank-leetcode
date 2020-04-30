package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestArrayManipulation(t *testing.T) {
	cases := []struct {
		result   []int
		expected []int
	}{
		{intersect([]int{1, 2, 2, 1}, []int{2, 2}), []int{2, 2}},
		{intersect([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}), []int{4, 9}},
		{intersect([]int{4, 9, 5, 9}, []int{9, 4, 9, 8, 4}), []int{4, 9, 9}},
	}
	for _, c := range cases {
		sort.Ints(c.result)
		if !reflect.DeepEqual(c.expected, c.result) {
			t.Errorf("Error got: %v, want: %v.", c.result, c.expected)
		}
	}
}
