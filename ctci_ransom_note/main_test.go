package main

import (
	"testing"
)

func TestCheckMagazine(t *testing.T) {
	cases := []struct {
		result   string
		expected string
	}{
		{
			checkMagazine(
				[]string{"give", "me", "one", "grand", "today", "night"},
				[]string{"give", "me", "one", "grand", "today"}), "Yes",
		},
		{
			checkMagazine(
				[]string{"two", "times", "three", "is", "not", "four"},
				[]string{"two", "times", "two", "is", "four"}), "No",
		},
		{
			checkMagazine(
				[]string{"ive", "got", "a", "lovely", "bunch", "of", "coconuts"},
				[]string{"ive", "got", "some", "coconuts"}), "No",
		},
	}
	for _, c := range cases {
		if c.expected != c.result {
			t.Errorf("Error got: %v, want: %v.", c.result, c.expected)
		}
	}
}
