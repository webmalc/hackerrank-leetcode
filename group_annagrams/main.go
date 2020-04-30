package main

import (
	"sort"
	"strings"
)

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

// https://leetcode.com/explore/interview/card/top-interview-questions-medium/103/array-and-strings/778/
func groupAnagrams(str []string) [][]string {
	table := make(map[string][]string)
	result := make([][]string, 0)
	for i, v := range str {
		s := sortString(v)
		table[s] = append(table[s], str[i])
	}
	for _, v := range table {
		result = append(result, v)
	}
	return result
}
