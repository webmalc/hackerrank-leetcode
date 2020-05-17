package main

import (
	"testing"
)

func TestSherlockAndAnagrams(t *testing.T) {
	cases := []struct {
		result   int
		expected int
	}{
		{
			trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}), 6,
		},
		{
			trap([]int{0, 1, 0, 99999, 1, 0, 1, 3, 2, 1, 2, 1}), 9,
		},
		{
			trap([]int{0, 1, 0, 0, 2, 0}), 2,
		},
		{
			trap([]int{0, 2, 0, 0, 2, 0}), 4,
		},
		{
			trap([]int{0, 2, 1, 0, 2, 0}), 3,
		},
		{
			trap([]int{0, 2, 1, 0, 2, 0, 1, 0, 2, 0}), 8,
		},
		{
			trap([]int{}), 0,
		},
		{
			trap([]int{1, 1, 1, 1}), 0,
		},
		{
			trap([]int{0, 0, 0, 1}), 0,
		},
	}
	for i, c := range cases {
		if c.expected != c.result {
			t.Errorf("%d: Error got: %v, want: %v.", i+1, c.result, c.expected)
		}
	}
}
