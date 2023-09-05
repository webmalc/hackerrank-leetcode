package main

import "fmt"

func longestCommonPrefix(strs []string) string {
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

func main() {
	nums := []string{"a"}

	r := longestCommonPrefix(nums)

	fmt.Println("-----------RESULT---------------")
	fmt.Println(r)
}
