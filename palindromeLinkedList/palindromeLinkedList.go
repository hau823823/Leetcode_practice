package palindromeLinkedList

import (
	"leetcode/node"
)

// brute
func IsPalindrome(head *node.ListNode) bool {
	ptr := head
	tmp := []int{}

	for ptr != nil {
		tmp = append(tmp, ptr.Val)
		ptr = ptr.Next
	}

	for i := 0; i < len(tmp); i++ {
		if tmp[i] != tmp[len(tmp)-i-1] {
			return false
		}
	}

	return true
}

// two pointer
func IsPalindromeTP(head *node.ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// reverse
	f := slow
	s := slow.Next
	f.Next = nil
	for s != nil {
		tmp := s.Next
		s.Next = f
		f = s
		s = tmp
	}

	for f != nil {
		v1 := head.Val
		v2 := f.Val
		if v1 != v2 {
			return false
		}
		head = head.Next
		f = f.Next
	}
	return true
}
