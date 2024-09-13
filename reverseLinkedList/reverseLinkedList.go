package reverseLinkedList

import (
	"fmt"
	"leetcode/node"
)

// false
func reverseList(head *node.ListNode) *node.ListNode {
	ptr := head
	last := ptr
	next := ptr.Next
	for next != nil {
		fmt.Println(ptr)
		ptr = ptr.Next
		next.Next = last
		last = ptr
		next = ptr.Next
	}
	return next
}

func ReverseList(head *node.ListNode) *node.ListNode {
    var last *node.ListNode = nil
    ptr := head

    for ptr != nil {
        next := ptr.Next // 保存当前节点的下一个节点
        ptr.Next = last  // 反转当前节点的指向
        last = ptr       // 更新 last 为当前节点
        ptr = next       // 移动到原链表的下一个节点
    }
    
    return last // 反转后的头节点
}
