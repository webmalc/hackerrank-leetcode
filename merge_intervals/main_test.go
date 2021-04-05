package mergeintervals

import (
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	cases := []struct {
		result   [][]int
		expected [][]int
	}{
		{merge([][]int{}), [][]int{}},
		{
			merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}),
			[][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			merge([][]int{{1, 3}, {3, 7}, {1, 2}, {8, 10}, {15, 18}}),
			[][]int{{1, 7}, {8, 10}, {15, 18}},
		},
		{
			merge([][]int{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}}),
			[][]int{{1, 10}},
		},
		{
			merge([][]int{{1, 3}, {10, 12}, {1, 5}, {8, 10}, {15, 18}}),
			[][]int{{1, 5}, {8, 12}, {15, 18}},
		},
		{
			merge([][]int{{1, 4}, {4, 5}}),
			[][]int{{1, 5}},
		},
	}
	for _, c := range cases {
		if !reflect.DeepEqual(c.expected, c.result) {
			t.Errorf("Error got: %v, want: %v.", c.result, c.expected)
		}
	}
}
