package canconstruct

// https://leetcode.com/problems/ransom-note
func CanConstruct(ransomNote string, magazine string) bool {
	letters := make(map[rune]int)
	for _, char := range magazine {
		letters[char]++
	}
	for _, char := range ransomNote {
		num, ok := letters[char]
		if !ok || num <= 0 {
			return false
		}
		letters[char]--
	}
	return true
}
