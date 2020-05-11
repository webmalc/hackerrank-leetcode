package main

import (
	"testing"
)

func TestSherlockAndAnagrams(t *testing.T) {
	cases := []struct {
		result   int
		expected int
	}{
		{
			sherlockAndAnagrams("xyyx"), 4,
		},
		{
			sherlockAndAnagrams("abba"), 4,
		},
		{
			sherlockAndAnagrams("abcd"), 0,
		},
		{
			sherlockAndAnagrams("ifailuhkqq"), 3,
		},
		{
			sherlockAndAnagrams("kkkk"), 10,
		},
		{
			sherlockAndAnagrams("cdcd"), 5,
		},
	}
	for _, c := range cases {
		if c.expected != c.result {
			t.Errorf("Error got: %v, want: %v.", c.result, c.expected)
		}
	}
}
