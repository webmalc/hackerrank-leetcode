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

// https://leetcode.com/problems/remove-duplicates-from-sorted-array/
func RemoveDuplicates(nums []int) int {
	i := 1
	for k, v := range nums {
		if k == 0 {
			continue
		}
		if nums[i-1] != v {
			nums[i] = v
			i++
		}
	}
	return i
}
