package mergeksortedlists

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	cases := []struct {
		result   []int
		expected []int
	}{
		// naive
		{
			linkedListToSlice(mergeKListsNaive(slicesToLinkedLists([][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}}))),
			[]int{1, 1, 2, 3, 4, 4, 5, 6},
		},
		{
			linkedListToSlice(mergeKListsNaive(slicesToLinkedLists([][]int{{5}, {1, 3, 4, 7}, {2, 6}}))),
			[]int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			linkedListToSlice(mergeKListsNaive(slicesToLinkedLists([][]int{}))),
			[]int{},
		},
		{
			linkedListToSlice(mergeKListsNaive(slicesToLinkedLists([][]int{{}}))),
			[]int{},
		},
		// priority queue
		{
			linkedListToSlice(mergeKLists(slicesToLinkedLists([][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}}))),
			[]int{1, 1, 2, 3, 4, 4, 5, 6},
		},
		{
			linkedListToSlice(mergeKLists(slicesToLinkedLists([][]int{{5}, {1, 3, 4, 7}, {2, 6}}))),
			[]int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			linkedListToSlice(mergeKLists(slicesToLinkedLists([][]int{}))),
			[]int{},
		},
		{
			linkedListToSlice(mergeKLists(slicesToLinkedLists([][]int{{}}))),
			[]int{},
		},
	}
	for _, c := range cases {
		if !reflect.DeepEqual(c.expected, c.result) {
			t.Errorf("Error got: %v, want: %v.", c.result, c.expected)
		}
	}
}
