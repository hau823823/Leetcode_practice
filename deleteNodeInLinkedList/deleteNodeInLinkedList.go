package deleteNodeInLinkedList

import "leetcode/node"

func DeleteNode(node *node.ListNode) {
	node.Val, node.Next = node.Next.Val, node.Next.Next
}

func DeleteNodeOP(node *node.ListNode) {
    *node = *node.Next
}

