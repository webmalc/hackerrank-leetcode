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
		// {twoSum([]int{2, 7, 11, 15}, 9), []int{0, 1}},
		// {twoSum([]int{2, 3, 0, 12, 3, 4, 9}, 21), []int{3, 6}},
		{twoSum([]int{3, 2, 4}, 6), []int{1, 2}},
	}
	for _, c := range cases {
		if !reflect.DeepEqual(c.expected, c.result) {
			t.Errorf("Error got: %v, want: %v.", c.result, c.expected)
		}
	}
}
