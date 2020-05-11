package main

import (
	"testing"
)

func TestTwoStrings(t *testing.T) {
	cases := []struct {
		result   string
		expected string
	}{
		{
			twoStrings("hello", "world"), "YES",
		},
		{
			twoStrings("hi", "world"), "NO",
		},
	}
	for _, c := range cases {
		if c.expected != c.result {
			t.Errorf("Error got: %v, want: %v.", c.result, c.expected)
		}
	}
}
