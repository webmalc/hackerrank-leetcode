package romantoint

func RomanToInt(s string) int {
	romanMap := map[rune]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	n := len(s)
	num := romanMap[rune(s[n-1])]
	for i := n - 2; i >= 0; i-- {
		if romanMap[rune(s[i])] >= romanMap[rune(s[i+1])] {
			num += romanMap[rune(s[i])]
		} else {
			num -= romanMap[rune(s[i])]
		}
	}
	return num
}
