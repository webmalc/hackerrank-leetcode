package luckbalance

import (
	"reflect"
	"testing"
)

func TestMinimumAbsoluteDifference(t *testing.T) {
	cases := []struct {
		result   int32
		expected int32
	}{
		{luckBalance(2, [][]int32{{5, 1}, {1, 1}, {4, 0}}), 10},
		{luckBalance(1, [][]int32{{5, 1}, {1, 1}, {4, 0}}), 8},
		{luckBalance(
			3, [][]int32{{5, 1}, {2, 1}, {1, 1}, {8, 1}, {10, 0}, {5, 0}}), 29,
		},
	}
	for _, c := range cases {
		if !reflect.DeepEqual(c.expected, c.result) {
			t.Errorf("Error got: %v, want: %v.", c.result, c.expected)
		}
	}
}
