package getminimumcost

import (
	"reflect"
	"testing"
)

func TestGetMinimumCost(t *testing.T) {
	cases := []struct {
		result   int
		expected int
	}{
		{getMinimumCost(3, []int{2, 5, 6}), 13},
		{getMinimumCost(2, []int{2, 5, 6}), 15},
		{getMinimumCost(3, []int{1, 3, 5, 7, 9}), 29},
	}
	for _, c := range cases {
		if !reflect.DeepEqual(c.expected, c.result) {
			t.Errorf("Error got: %v, want: %v.", c.result, c.expected)
		}
	}
}
