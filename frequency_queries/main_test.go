package main

import (
	"reflect"
	"testing"
)

func TestFreqQuery(t *testing.T) {
	cases := []struct {
		result   []int
		expected []int
	}{
		{
			freqQuery(
				[][]int{{1, 5}, {1, 6}, {3, 2}, {1, 10}, {1, 10}, {1, 6}, {2, 5}, {3, 2}},
			),
			[]int{0, 1},
		},
		{
			freqQuery(
				[][]int{{3, 4}, {2, 1003}, {1, 16}, {3, 1}},
			),
			[]int{0, 1},
		},
		{
			freqQuery(
				[][]int{{1, 3}, {2, 3}, {3, 2}, {1, 4}, {1, 5}, {1, 5}, {1, 4}, {3, 2}, {2, 4}, {3, 2}},
			),
			[]int{0, 1, 1},
		},
	}
	for _, c := range cases {
		if !reflect.DeepEqual(c.expected, c.result) {
			t.Errorf("Error got: %v, want: %v.", c.result, c.expected)
		}
	}
}
