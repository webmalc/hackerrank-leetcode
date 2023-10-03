package strstr

// https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/
func StrStr(haystack string, needle string) int {
	i, r := 0, 0
	for i < len(haystack) {
		if haystack[i] == needle[r] {
			i++
			r++
			if r == len(needle) {
				return i - r
			}
		} else {
			i = i - r + 1
			r = 0
		}
	}
	return -1
}
