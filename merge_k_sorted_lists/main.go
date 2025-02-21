package mergeksortedlists

import (
	"container/heap"
	"errors"
	"math"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// An Item for priority queue
type Item struct {
	value    *ListNode
	priority int
	index    int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func slicesToLinkedLists(slices [][]int) []*ListNode {
	result := []*ListNode{}
	for _, slice := range slices {
		list, err := sliceToLinkedList(slice)
		if err == nil {
			result = append(result, list)
		}
	}

	return result
}

func sliceToLinkedList(slice []int) (*ListNode, error) {
	if len(slice) == 0 {
		return &ListNode{}, errors.New("slice is empty")
	}
	current := &ListNode{Val: slice[0]}
	head := current
	for _, value := range slice[1:] {
		next := &ListNode{Val: value}
		current.Next = next
		current = next
	}

	return head, nil
}

func linkedListToSlice(list *ListNode) []int {
	result := []int{}
	current := list
	for current != nil {
		result = append(result, current.Val)
		current = current.Next
	}

	return result
}

func findSmallestNode(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	min := math.MaxInt
	var minIndex = -1
	for i := range lists {
		if lists[i] == nil {
			continue
		}
		if lists[i].Val < min {
			min = lists[i].Val
			minIndex = i
		}
	}
	if minIndex != -1 {
		lists[minIndex] = lists[minIndex].Next
	} else {
		return nil
	}

	return &ListNode{Val: min}
}

// https://leetcode.com/problems/merge-k-sorted-lists/
func mergeKListsNaive(lists []*ListNode) *ListNode {
	fakeHead := &ListNode{}
	current := fakeHead
	min := findSmallestNode(lists)
	for min != nil && min.Val != math.MaxInt {
		current.Next = min
		current = min
		min = findSmallestNode(lists)
	}
	return fakeHead.Next
}

// https://leetcode.com/problems/merge-k-sorted-lists/
func mergeKLists(lists []*ListNode) *ListNode {
	pq := &PriorityQueue{}
	heap.Init(pq)
	fakeHead := &ListNode{}
	current := fakeHead

	for i := range lists {
		if lists[i] == nil {
			continue
		}
		item := &Item{value: lists[i], priority: lists[i].Val}
		heap.Push(pq, item)
	}
	for pq.Len() > 0 {
		item := heap.Pop(pq).(*Item)
		current.Next = item.value
		current = item.value
		next := item.value.Next
		if item.value.Next != nil {
			item = &Item{value: next, priority: next.Val}
			heap.Push(pq, item)
		}
	}

	return fakeHead.Next
}
