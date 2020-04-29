package main

//https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/546/
func twoSum(nums []int, target int) []int {
	hash := make(map[int]int)
	for i, v := range nums {
		hash[target-v] = i
	}
	for i, v := range nums {
		if h, ok := hash[v]; ok && i != h {
			if i < h {
				return []int{i, h}
			}
			return []int{h, i}
		}
	}
	return []int{}
}
