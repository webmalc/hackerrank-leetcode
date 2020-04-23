package main

import (
	"testing"
)

func TestMinimumBribes(t *testing.T) {
	cases := []struct {
		result   string
		expected string
	}{
		{minimumBribes([]int32{2, 1, 5, 3, 4}), "3"},
		{minimumBribes([]int32{2, 5, 1, 3, 4}), "Too chaotic"},
		{minimumBribes([]int32{1, 2, 5, 3, 7, 8, 6, 4}), "7"},
	}
	for _, c := range cases {
		if c.expected != c.result {
			t.Errorf("Error got: %s, want: %s.", c.result, c.expected)
		}
	}
}
