package oddEvenLinkedList

import (
	"leetcode/node"
)

func OddEvenList(head *node.ListNode) *node.ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}
	odd := head
	even := head.Next
	tmp := head.Next.Next
	tmpOdd := odd
	tmpEven := even
	count := 0

	for tmp != nil {
		if count%2 == 0 {
			tmpOdd.Next = tmp
			tmpOdd = tmpOdd.Next
		} else {
			tmpEven.Next = tmp
			tmpEven = tmpEven.Next
		}
		tmp = tmp.Next
		count++
	}
	tmpOdd.Next = even
    tmpEven.Next = nil // 這邊是重點
	return odd
}
