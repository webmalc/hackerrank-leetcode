package main

func makeHash(nums []int) map[int]int {
	hash := make(map[int]int)
	for _, i := range nums {
		hash[i]++
	}
	return hash
}

//https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/674/
func intersect(nums1, nums2 []int) []int {
	hash := makeHash(nums1)
	result := []int{}
	for _, v := range nums2 {
		if hash[v] > 0 {
			result = append(result, v)
			hash[v]--
		}
	}
	return result
}
