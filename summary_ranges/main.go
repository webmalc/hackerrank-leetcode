package summaryranges

import "fmt"

// Format the couple of ints
func format(start, end int) string {
	if start == end {
		return fmt.Sprintf("%v", start)
	}
	return fmt.Sprintf("%v->%v", start, end)
}

// https://leetcode.com/problems/summary-ranges/
func summaryRanges(nums []int) []string {
	result := []string{}
	if len(nums) == 0 {
		return result
	}
	start := nums[0]
	end := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1]+1 == nums[i] {
			end = nums[i]
		} else {
			result = append(result, format(start, end))
			start = nums[i]
			end = nums[i]
		}
	}
	result = append(result, format(start, end))
	return result
}
