package main

import (
	"testing"
)

func TestActivityNotifications(t *testing.T) {
	cases := []struct {
		result   int
		expected int
	}{
		{
			countInversions([]int{1, 1, 1, 2, 2}), 0,
		},
		{
			countInversions([]int{2, 1, 3, 1, 2}), 4,
		},
	}
	for _, c := range cases {
		if c.result != c.expected {
			t.Errorf("Error got: %v, want: %v.", c.result, c.expected)
		}
	}
}
