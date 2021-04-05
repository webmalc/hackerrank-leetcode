package mergeintervals

import (
	"sort"
)

// Get the max between a and b
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// https://leetcode.com/problems/merge-intervals/
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	if len(intervals) == 0 {
		return intervals
	}
	result := append([][]int{}, intervals[0])
	for i := 1; i < len(intervals); i++ {
		current := intervals[i]
		lastResultIndex := len(result) - 1
		prev := result[lastResultIndex]
		if prev[1] < current[0] {
			result = append(result, current)
		} else {
			result[lastResultIndex][1] = max(prev[1], current[1])
		}
	}
	return result
}
