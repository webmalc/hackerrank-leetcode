package main

import (
	"testing"
)

func TestMaximumToys(t *testing.T) {
	cases := []struct {
		result   int
		expected int
	}{
		{maximumToys([]int{1, 12, 5, 111, 200, 1000, 10}, 50), 4},
		{maximumToys([]int{3, 7, 2, 9, 4}, 15), 3},
	}
	for _, c := range cases {
		if c.expected != c.result {
			t.Errorf("Error got: %d, want: %d.", c.result, c.expected)
		}
	}
}
