package main

import (
	"testing"
)

func TestMakeAnagram(t *testing.T) {
	cases := []struct {
		result   string
		expected string
	}{
		{isValid("aabbcd"), "NO"},
		{isValid("aabbccddeefghi"), "NO"},
		{isValid("abcdefghhgfedecba"), "YES"},
		{isValid("aaaabbcc"), "NO"},
	}
	for _, c := range cases {
		if c.expected != c.result {
			t.Errorf("Error got: %s, want: %s.", c.result, c.expected)
		}
	}
}
