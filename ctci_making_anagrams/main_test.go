package main

import (
	"testing"
)

func TestMakeAnagram(t *testing.T) {
	cases := []struct {
		result   int
		expected int
	}{
		{makeAnagram("cde", "abc"), 4},
		{makeAnagram("abc", "abcdf"), 2},
		{makeAnagram("showman", "woman"), 2},
		{makeAnagram("fcrxzwscanmligyxyvym", "jxwtrhvujlmrpdoqbisbwhmgpmeoke"), 30},
	}
	for _, c := range cases {
		if c.expected != c.result {
			t.Errorf("Error got: %d, want: %d.", c.result, c.expected)
		}
	}
}
