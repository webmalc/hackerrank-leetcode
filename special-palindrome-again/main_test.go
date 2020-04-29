package main

import (
	"testing"
)

func TestMakeAnagram(t *testing.T) {
	cases := []struct {
		result   int64
		expected int64
	}{
		{substrCount(5, "asasd"), 7},
		{substrCount(7, "abcbaba"), 10},
		{substrCount(4, "aaaa"), 10},
	}
	for _, c := range cases {
		if c.expected != c.result {
			t.Errorf("Error got: %d, want: %d.", c.result, c.expected)
		}
	}
}
