package main

// https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/549/
func singleNumber(nums []int) int {
	hash := make(map[int]int)
	for _, i := range nums {
		hash[i]++
	}
	for i, c := range hash {
		if c == 1 {
			return i
		}
	}
	return 0
}
