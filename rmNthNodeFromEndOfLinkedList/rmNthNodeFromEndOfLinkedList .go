package rmNthNodeFromEndOfLinkedList

import (
	"fmt"
	"leetcode/node"
)

// brute
// will be wrong
func RemoveNthFromEnd(head *node.ListNode, n int) *node.ListNode {
	tmp := head
	count := 0
	for tmp != nil {
		tmp = tmp.Next
		count++
	}
	if count == 1 {
		return nil
	}

	tmp = head
	step := count - n - 1

	

	for step > 0 {
		tmp = tmp.Next
		step--
	}

	fmt.Println(tmp.Next)
	fmt.Println(tmp.Next.Next)
	tmp.Next = tmp.Next.Next

	return head
}

// two pointer
func RemoveNthFromEnd2PT(head *node.ListNode, n int) *node.ListNode {
	tmp := &node.ListNode{Next: head}
	fast, slow := tmp, tmp
	skip := n

	for fast.Next != nil {
		fast = fast.Next
		if skip == 0 {
			slow = slow.Next
		} else {
			skip --
		}
	}
	
	slow.Next = slow.Next.Next

	return tmp.Next
}