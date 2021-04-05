package summaryranges

import (
	"reflect"
	"testing"
)

func TestSummaryRanges(t *testing.T) {
	cases := []struct {
		result   []string
		expected []string
	}{
		{summaryRanges([]int{}), []string{}},
		{summaryRanges([]int{-1}), []string{"-1"}},
		{summaryRanges([]int{0}), []string{"0"}},
		{
			summaryRanges([]int{0, 1, 2, 4, 5, 7}),
			[]string{"0->2", "4->5", "7"},
		},
		{
			summaryRanges([]int{0, 2, 3, 4, 6, 8, 9}),
			[]string{"0", "2->4", "6", "8->9"},
		},
	}
	for _, c := range cases {
		if !reflect.DeepEqual(c.expected, c.result) {
			t.Errorf("Error got: %v, want: %v.", c.result, c.expected)
		}
	}
}
