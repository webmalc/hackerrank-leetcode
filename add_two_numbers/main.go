package addtwonumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

// Makes a linked list from a number
func makeList(nums []int) *ListNode {
	if len(nums) == 0 {
		return &ListNode{}
	}
	first := ListNode{Val: nums[0]}
	head := &first
	tail := &first
	for _, v := range nums[1:] {
		tail.Next = &ListNode{Val: v}
		tail = tail.Next
	}
	return head
}

func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	current1 := l1
	current2 := l2
	memory, num, sum := 0, 0, 0
	result := []int{}
	for current1 != nil || current2 != nil {
		if current1 != nil && current2 != nil {
			sum = current1.Val + current2.Val
			current1 = current1.Next
			current2 = current2.Next
		} else if current1 != nil {
			sum = current1.Val
			current1 = current1.Next
		} else {
			sum = current2.Val
			current2 = current2.Next
		}
		sum += memory
		num = sum % 10
		memory = sum / 10
		result = append(result, num)
	}
	if memory != 0 {
		result = append(result, memory)
	}
	return makeList(result)
}
