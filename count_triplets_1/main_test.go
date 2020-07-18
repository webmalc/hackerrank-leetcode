package main

import (
	"testing"
)

func TestCountTriplets(t *testing.T) {
	cases := []struct {
		result   int
		expected int
	}{
		{countTriplets([]int{1, 2, 2, 4, 5}, 2), 2},
		{countTriplets([]int{1, 3, 9, 9, 27, 81}, 3), 6},
		{countTriplets([]int{1, 5, 5, 25, 125}, 5), 4},
		{countTriplets([]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 1), 161700},
	}
	for _, c := range cases {
		if c.expected != c.result {
			t.Errorf("Error got: %d, want: %d.", c.result, c.expected)
		}
	}
}
