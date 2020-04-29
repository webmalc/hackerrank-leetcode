package main

import (
	"testing"
)

func TestArrayManipulation(t *testing.T) {
	cases := []struct {
		result   int
		expected int
	}{
		{singleNumber([]int{2, 2, 1}), 1},
		{singleNumber([]int{4, 1, 2, 1, 2}), 4},
	}
	for _, c := range cases {
		if c.expected != c.result {
			t.Errorf("Error got: %d, want: %d.", c.result, c.expected)
		}
	}
}
