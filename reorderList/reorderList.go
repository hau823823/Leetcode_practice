package reorderList

import (
	"leetcode/node"
)

func ReorderList(head *node.ListNode) {
	ptr := head
	mid := ptr
	for ptr != nil && ptr.Next != nil {
		ptr = ptr.Next.Next
		mid = mid.Next
	}

	// reverse
	ptr = mid.Next
	var last *node.ListNode = nil
	for ptr != nil {
		next := ptr.Next
		ptr.Next = last
		last = ptr
		ptr = next
	}
	mid.Next = nil

	// merge
	ptr1, ptr2 := head, last
	for ptr2 != nil {
		tmp1, tmp2 := ptr1.Next, ptr2.Next

        ptr1.Next = ptr2
        ptr1 = tmp1

        ptr2.Next = ptr1
        ptr2 = tmp2
	}
}
