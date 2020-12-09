package minimumabsolutedifferencearray

import (
	"reflect"
	"testing"
)

func TestMinimumAbsoluteDifference(t *testing.T) {
	cases := []struct {
		result   int
		expected int
	}{
		{minimumAbsoluteDifference([]int{3, -7, 0}), 3},
		{minimumAbsoluteDifference(
			[]int{-59, -36, -13, 1, -53, -92, -2, -96, -54, 75}), 1,
		},
		{minimumAbsoluteDifference([]int{1, -3, 71, 68, 17}), 3},
	}
	for _, c := range cases {
		if !reflect.DeepEqual(c.expected, c.result) {
			t.Errorf("Error got: %v, want: %v.", c.result, c.expected)
		}
	}
}
