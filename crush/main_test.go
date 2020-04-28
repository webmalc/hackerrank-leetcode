package main

import (
	"testing"
)

func TestArrayManipulation(t *testing.T) {
	cases := []struct {
		result   int64
		expected int64
	}{
		{arrayManipulation(5, [][]int32{{1, 2, 100}, {2, 5, 100}, {3, 4, 100}}), 200},
		{arrayManipulation(10, [][]int32{{1, 5, 3}, {4, 8, 7}, {6, 9, 1}}), 10},
	}
	for _, c := range cases {
		if c.expected != c.result {
			t.Errorf("Error got: %d, want: %d.", c.result, c.expected)
		}
	}
}
