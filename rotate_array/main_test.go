package main

import (
	"reflect"
	"testing"
)

func TestArrayManipulation(t *testing.T) {
	cases := []struct {
		result   []int
		expected []int
	}{
		{rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3), []int{5, 6, 7, 1, 2, 3, 4}},
		{rotate([]int{-1, -100, 3, 99}, 2), []int{3, 99, -1, -100}},
	}
	for _, c := range cases {
		if !reflect.DeepEqual(c.expected, c.result) {
			t.Errorf("Error got: %v, want: %v.", c.result, c.expected)
		}
	}
}
