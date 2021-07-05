package stringcompression

import (
	"strconv"
)

// https://leetcode.com/problems/string-compression/
func compress(chars []byte) int {
	sum := 0
	max := len(chars)
	result := make([]byte, 0)
	for i, current := range chars {
		sum++
		if i == max-1 || current != chars[i+1] {
			if sum == 1 {
				result = append(result, current)
			} else if sum > 1 {
				result = append(result, current)
				result = append(result, []byte(strconv.Itoa(sum))...)
			}
			sum = 0
		}
	}
	copy(chars, result)
	return len(result)
}
