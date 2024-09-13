package merge2SortedLists

import "leetcode/node"

func MergeTwoLists(list1 *node.ListNode, list2 *node.ListNode) *node.ListNode {

	head := &node.ListNode{}
	tmp := head
    ptr1, ptr2 := list1, list2
	
	for ptr1 != nil || ptr2 != nil {
		if ptr1 != nil && ptr2 != nil {
			if ptr1.Val < ptr2.Val{
				tmp.Next = ptr1
				ptr1 = ptr1.Next
			} else {
				tmp.Next = ptr2
				ptr2 = ptr2.Next
			}
		} else if ptr1 != nil {
			tmp.Next = ptr1
			ptr1 = ptr1.Next
		} else {
			tmp.Next = ptr2
			ptr2 = ptr2.Next
		}
		tmp = tmp.Next
	}

	return head.Next
}

func MergeTwoListsOP(list1 *node.ListNode, list2 *node.ListNode) *node.ListNode {
    head := &node.ListNode{}
	tmp := head
    ptr1, ptr2 := list1, list2
	
	for ptr1 != nil && ptr2 != nil {
		if ptr1.Val < ptr2.Val{
			tmp.Next = ptr1
			ptr1 = ptr1.Next
		} else {
			tmp.Next = ptr2
			ptr2 = ptr2.Next
		}
			
		tmp = tmp.Next
	}
    if ptr1 != nil {
        tmp.Next = ptr1
    }
    if ptr2 != nil {
        tmp.Next = ptr2
    }

	return head.Next
}