package main

import (
	"testing"
)

func TestMakeAnagram(t *testing.T) {
	cases := []struct {
		result   int
		expected int
	}{
		{alternatingCharacters("AAAA"), 3},
		{alternatingCharacters("BBBBB"), 4},
		{alternatingCharacters("ABABABAB"), 0},
		{alternatingCharacters("BABABA"), 0},
		{alternatingCharacters("AAABBB"), 4},
	}
	for _, c := range cases {
		if c.expected != c.result {
			t.Errorf("Error got: %d, want: %d.", c.result, c.expected)
		}
	}
}
