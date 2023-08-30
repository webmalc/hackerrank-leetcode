package mergesortedarray

// https://leetcode.com/problems/merge-sorted-array/
func merge(nums1 []int, m int, nums2 []int, n int) []int {
	l := m + n - 1
	for m > 0 && n > 0 {
		if nums1[m-1] > nums2[n-1] {
			nums1[l] = nums1[m-1]
			m--
		} else {
			nums1[l] = nums2[n-1]
			n--
		}
		l--
	}
	for n > 0 {
		nums1[l] = nums2[n-1]
		n--
		l--
	}
	return nums1
}
