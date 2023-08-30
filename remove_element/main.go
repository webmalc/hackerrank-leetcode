package removeelement

// https://leetcode.com/problems/remove-element/
func RemoveElement(nums []int, val int) int {
	k := 0
	for _, x := range nums {
		if x != val {
			nums[k] = x
			k++
		}
	}
	return k
}
