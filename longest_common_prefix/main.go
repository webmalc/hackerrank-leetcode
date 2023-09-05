package removeelement

// https://leetcode.com/problems/longest-common-prefix/
func LongestCommonPrefix(strs []string) string {
	for i, ch := range strs[0] {
		for k, word := range strs {
			if k == 0 {
				continue
			}
			if len(word) <= i || ch != rune(word[i]) {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}
