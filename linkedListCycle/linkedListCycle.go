package linkedListCycle

import "leetcode/node"

func HasCycle(head *node.ListNode) bool {
	fast, slow := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			return true
		}
	}

	return false
}