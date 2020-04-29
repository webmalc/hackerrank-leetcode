package main

// https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/646/
func rotate(nums []int, k int) []int {
	l := len(nums)
	c := k % l
	copy(nums, append(nums[l-c:], nums[:l-c]...))

	return nums
}
