package addtwonumbers

import (
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	cases := []struct {
		result   *ListNode
		expected *ListNode
	}{
		{addTwoNumbers(
			makeList([]int{4, 3}),
			makeList([]int{3, 6, 4})),
			makeList([]int{7, 9, 4}),
		},
		{addTwoNumbers(
			makeList([]int{2, 4, 3}),
			makeList([]int{5, 6, 4})),
			makeList([]int{7, 0, 8}),
		},
		{addTwoNumbers(
			makeList([]int{0}),
			makeList([]int{0})),
			makeList([]int{0}),
		},
		{addTwoNumbers(
			makeList([]int{9, 9, 9, 9, 9, 9, 9}),
			makeList([]int{9, 9, 9, 9})),
			makeList([]int{8, 9, 9, 9, 0, 0, 0, 1}),
		},
	}
	for _, c := range cases {
		if !reflect.DeepEqual(c.expected, c.result) {
			t.Errorf("Error got: %v, want: %v.", c.result, c.expected)
		}
	}
}
