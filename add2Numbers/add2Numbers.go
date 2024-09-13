package add2Numbers

import (
	"leetcode/node"
)

func AddTwoNumbers(l1 *node.ListNode, l2 *node.ListNode) *node.ListNode {
	res := new(node.ListNode)
	tmp := res
	tmp1 := l1
	tmp2 := l2
	carry := 0

	for tmp1 != nil || tmp2 != nil || carry > 0 {
		sum := carry
		if tmp1 != nil {
			sum += tmp1.Val
			tmp1 = tmp1.Next
		}
		if tmp2 != nil {
			sum += tmp2.Val
			tmp2 = tmp2.Next
		}

		carry = sum / 10
		tmp.Val = sum % 10

		if tmp1 != nil || tmp2 != nil || carry > 0 {
			tmp.Next = new(node.ListNode)
			tmp = tmp.Next
		}
	}

	return res
}