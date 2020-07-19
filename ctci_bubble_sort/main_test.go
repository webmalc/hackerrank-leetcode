package main

import (
	"testing"
)

func TestCountSwaps(t *testing.T) {
	cases := []struct {
		result   [3]int
		expected [3]int
	}{
		{countSwaps([]int{1, 2, 3}), [3]int{0, 1, 3}},
		{countSwaps([]int{3, 2, 1}), [3]int{3, 1, 3}},
	}
	for _, c := range cases {
		if c.expected != c.result {
			t.Errorf("Error got: %d, want: %d.", c.result, c.expected)
		}
	}
}
