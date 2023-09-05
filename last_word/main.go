package removeelement

// https://leetcode.com/problems/length-of-last-word/
func LengthOfLastWord(s string) int {
	max := 0
	next := true
	for _, char := range s {
		if char != ' ' && next {
			max = 0
			next = false
		}
		if char != ' ' {
			max++
		} else {
			next = true
		}
	}
	return max
}
