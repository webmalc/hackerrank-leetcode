package angrychildren

import (
	"reflect"
	"testing"
)

func TestGetMinimumCost(t *testing.T) {
	cases := []struct {
		result   int
		expected int
	}{
		{maxMin(2, []int{1, 4, 7, 2}), 1},
		{maxMin(3, []int{10, 100, 300, 200, 1000, 20, 30}), 20},
		{maxMin(4, []int{1, 2, 3, 4, 10, 20, 30, 40, 100, 200}), 3},
		{maxMin(2, []int{1, 2, 1, 2, 1}), 0},
	}
	for _, c := range cases {
		if !reflect.DeepEqual(c.expected, c.result) {
			t.Errorf("Error got: %v, want: %v.", c.result, c.expected)
		}
	}
}
