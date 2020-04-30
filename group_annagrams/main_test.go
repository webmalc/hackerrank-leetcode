package main

import (
	"reflect"
	"testing"
)

func TestArrayManipulation(t *testing.T) {
	cases := []struct {
		result   [][]string
		expected [][]string
	}{
		{
			groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}),
			[][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}},
		},
	}
	for _, c := range cases {
		if !reflect.DeepEqual(c.expected, c.result) {
			t.Errorf("Error got: %v, want: %v.", c.result, c.expected)
		}
	}
}
