package leetcodemisc

import (
	"sort"
	"strconv"
	"strings"
)

func LongestConsecutive(nums []int) int {
	numsMap := make(map[int]int)
	if len(nums) == 0 {
		return 0
	}
	for _, num := range nums {
		numsMap[num] = 0
	}
	count := 0
	for i := range numsMap {
		_, ok := numsMap[i-1]
		j := i
		if !ok {
			for {
				_, ok = numsMap[i+1]
				if !ok {
					if count < i-j {
						count = i - j
					}
					break
				} else {
					i++
				}
			}
		}
	}

	return count + 1
}

func isP(s string) bool {
	l := len(s)
	for i := 0; i < l/2; i++ {
		if s[i] != s[l-i-1] {
			return false
		}
	}
	return true
}

func LongestPalindrome(s string) string {
	l := len(s)
	result := s[0:1]
	for i := 0; i < l; i++ {
		for k := i + 1; k < l; k++ {
			if s[i] == s[k] {
				sub := s[i : k+1]
				if isP(sub) {
					if len(sub) > len(result) {
						result = sub
					}
				}
			}
		}
	}
	return result
}

func LengthOfLongestSubstring(s string) int {
	result := 0

	l := len(s)
	for i := 0; i < l; i++ {
		ch := make([]bool, 256)
		for k := i; k < l; k++ {
			if ch[s[k]] {
				break
			} else {
				ch[s[k]] = true
				g := k - i + 1
				if g > result {
					result = g
				}
			}
		}
		ch[s[i]] = false
	}
	return result
}

func GetConcatenation(nums []int) []int {
	return append(nums, nums...)
}

// https://leetcode.com/problems/baseball-game/
func CalPoints(operations []string) int {
	nums := []int{}
	sum := 0
	for _, o := range operations {
		l := len(nums) - 1
		switch o {
		case "C":
			nums = nums[:l]
		case "D":
			nums = append(nums, nums[l]*2)
		case "+":
			nums = append(nums, nums[l]+nums[l-1])
		default:
			num, _ := strconv.Atoi(o)
			nums = append(nums, num)
		}
	}
	for _, v := range nums {
		sum += v
	}

	return sum
}

// https://leetcode.com/problems/valid-parentheses/
func IsValid(s string) bool {
	stack := []string{}
	for _, r := range s {
		b := string(r)
		if b == "(" || b == "[" || b == "{" {
			stack = append(stack, b)
		} else {
			l := len(stack)
			if l == 0 {
				return false
			}
			switch b {
			case ")":
				if stack[l-1] != "(" {
					return false
				}
			case "]":
				if stack[l-1] != "[" {
					return false
				}
			case "}":
				if stack[l-1] != "{" {
					return false
				}
			}
			stack = stack[:l-1]
		}
	}
	return len(stack) == 0
}

// https://leetcode.com/problems/min-stack/
type MinStack struct {
	mins []int
	nums []int
}

func Constructor() MinStack {
	return MinStack{mins: []int{}, nums: []int{}}
}

func (s *MinStack) Push(val int) {
	if len(s.mins) == 0 || val <= s.GetMin() {
		s.mins = append(s.mins, val)
	}
	s.nums = append(s.nums, val)
}

func (s *MinStack) Pop() {
	i := len(s.nums) - 1
	if i <= 0 {
		s.nums = []int{}
		s.mins = []int{}
	} else {
		if s.Top() == s.GetMin() {
			k := len(s.mins) - 1
			if k <= 0 {
				s.mins = []int{}
			} else {
				s.mins = s.mins[:k]
			}
		}
		s.nums = s.nums[:i]
	}
}

func (s *MinStack) Top() int {
	i := len(s.nums) - 1
	if i < 0 {
		return 0
	}
	return s.nums[i]
}

func (s *MinStack) GetMin() int {
	i := len(s.mins) - 1
	if i < 0 {
		return 0
	}
	return s.mins[i]
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.com/problems/reverse-linked-list/
func ReverseList(head *ListNode) *ListNode {
	var p, c, n *ListNode
	c = head
	for c != nil {
		n = c.Next
		c.Next = p
		p = c
		c = n
	}
	return p
}

// https://leetcode.com/problems/merge-two-sorted-lists/description/
func MergeTwoLists(list1, list2 *ListNode) *ListNode {
	nodes := []*ListNode{}

	c := list1
	for c != nil {
		nodes = append(nodes, c)
		c = c.Next
	}

	c = list2
	for c != nil {
		nodes = append(nodes, c)
		c = c.Next
	}

	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].Val < nodes[j].Val
	})
	l := len(nodes)
	if l == 0 {
		return nil
	}
	l--
	for i := range nodes {
		if i < l {
			nodes[i].Next = nodes[i+1]
		} else {
			nodes[i].Next = nil
		}
	}

	return nodes[0]
}

func MergeTwoLists2(list1, list2 *ListNode) *ListNode {
	o := list1
	t := list2
	var c *ListNode
	var h *ListNode
	for o != nil || t != nil {
		if o == nil {
			if h == nil {
				h = t
				c = h
			} else {
				c.Next = t
				c = c.Next
			}
			t = t.Next
		} else if t == nil {
			if h == nil {
				h = o
				c = h
			} else {
				c.Next = o
				c = c.Next
			}
			o = o.Next
		} else {
			if o.Val < t.Val {
				if h == nil {
					h = o
					c = h
				} else {
					c.Next = o
					c = c.Next
				}
				o = o.Next
			} else {
				if h == nil {
					h = t
					c = h
				} else {
					c.Next = t
					c = c.Next
				}
				t = t.Next
			}
		}
	}
	return h
}

// https://leetcode.com/problems/valid-anagram
func IsAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	letters := make(map[rune]int)
	for _, char := range s {
		letters[char]++
	}
	for _, char := range t {
		v, ok := letters[char]
		if !ok || v == 0 {
			return false
		}
		letters[char]--
	}
	return true
}

// https://leetcode.com/problems/isomorphic-strings/
func IsIsomorphic(s, t string) bool {
	left := make(map[string]string)
	right := make(map[string]string)
	for i, char := range s {
		mapChar, ok := left[string(char)]
		if !ok {
			left[string(char)] = string(t[i])
			right[string(t[i])] = string(char)
		} else if mapChar != string(t[i]) {
			return false
		}
	}
	return true && len(left) == len(right)
}

// https://leetcode.com/problems/word-pattern/
func WordPattern(pattern, s string) bool {
	t := strings.Split(s, " ")
	left := make(map[string]string)
	right := make(map[string]string)
	if len(t) != len(pattern) {
		return false
	}
	for i, char := range pattern {
		mapChar, ok := left[string(char)]
		if !ok {
			left[string(char)] = t[i]
			_, ok = right[t[i]]
			if ok && right[t[i]] != string(char) {
				return false
			}
			right[t[i]] = string(char)
		} else if mapChar != string(t[i]) {
			return false
		}
	}

	return true
}
