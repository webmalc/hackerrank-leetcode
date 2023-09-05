package removeelement

// https://leetcode.com/problems/majority-element/
func MajorityElement(nums []int) int {
	table := make(map[int]int)
	max := table[nums[0]]
	res := nums[0]
	for _, v := range nums {
		table[v]++
		if table[v] > max {
			max = table[v]
			res = v
		}
	}
	return res
}
