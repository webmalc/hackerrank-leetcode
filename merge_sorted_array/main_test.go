package mergesortedarray

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	cases := []struct {
		result   []int
		expected []int
	}{
		{merge(
			[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3), []int{1, 2, 2, 3, 5, 6},
		},
		{merge(
			[]int{2, 5, 6, 0, 0, 0}, 3, []int{1, 2, 3}, 3), []int{1, 2, 2, 3, 5, 6},
		},
		{merge(
			[]int{0, 0, 0, 0, 0}, 0, []int{1, 2, 3, 4, 5}, 5), []int{1, 2, 3, 4, 5},
		},
		{merge(
			[]int{1}, 1, []int{}, 0), []int{1},
		},
		{merge(
			[]int{0}, 0, []int{1}, 1), []int{1},
		},
	}
	for _, c := range cases {
		if !reflect.DeepEqual(c.expected, c.result) {
			t.Errorf("Error got: %v, want: %v.", c.result, c.expected)
		}
	}
}
