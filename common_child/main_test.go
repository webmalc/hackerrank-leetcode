package main

import (
	"testing"
)

func TestMakeAnagram(t *testing.T) {
	cases := []struct {
		result   int
		expected int
	}{
		{commonChild("HARRY", "SALLY"), 2}, // AY
		{commonChild("AA", "BB"), 0},
		{commonChild("SHINCHAN", "NOHARAAA"), 3}, // NHA
		{commonChild("ABCDEF", "FBDAMN"), 2},     // BD
	}
	for _, c := range cases {
		if c.expected != c.result {
			t.Errorf("Error got: %d, want: %d.", c.result, c.expected)
		}
	}
}
