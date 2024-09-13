package intersection2LinkedList

import (
	"fmt"
	"leetcode/node"
)

func Generate(listA, listB []int, skipA, skipB int) (headA, headB *node.ListNode) {
	headA = &node.ListNode{Val: listA[0]}
	headB = &node.ListNode{Val: listB[0]}
	tmpA, tmpB := headA, headB

	for i := 1; i < skipA; i++ {
		tmpA.Next = &node.ListNode{Val: listA[i]}
		tmpA = tmpA.Next
	}
	for i := 1; i < skipB; i++ {
		tmpB.Next = &node.ListNode{Val: listB[i]}
		tmpB = tmpB.Next
	}

	if (len(listA)-skipA == 0) || (len(listB)-skipB == 0) {
		return headA, headB
	}

	for i := skipA; i < len(listA); i++ {
		if i == skipA {
			tmpB.Next = &node.ListNode{Val: listA[i]}
			tmpA.Next = tmpB.Next
			tmpA = tmpA.Next
		} else {
			tmpA.Next = &node.ListNode{Val: listA[i]}
			tmpA = tmpA.Next
		}
	}

	return headA, headB
}

// origin way
func GetIntersectionNode(headA, headB *node.ListNode) *node.ListNode {
	ptrA, ptrB := headA, headB
	countA, countB := 0, 0
	for ptrA != nil {
		countA++
		ptrA = ptrA.Next
	}
	for ptrB != nil {
		countB++
		ptrB = ptrB.Next
	}

	tmpA, tmpB := headA, headB
	if countB > countA {
		countA, countB = countB, countA
		tmpA, tmpB = tmpB, tmpA
	}

	skip := countA - countB
	for tmpA != nil && tmpB != nil {
		if tmpA == tmpB {
			return tmpA
		}
		tmpA = tmpA.Next
		if skip == 0 {
			tmpB = tmpB.Next
		} else {
			skip--
		}
	}

	return nil
}

// optimize
func GetIntersectionNodeOP(headA, headB *node.ListNode) *node.ListNode {
	tmpA, tmpB := headA, headB
	count := 1

	for tmpA != tmpB {
		fmt.Println("count:", count)
		if tmpA == nil {
			tmpA = headB
		} else {
			tmpA = tmpA.Next
		}
		fmt.Println("tmpA:", tmpA)
		if tmpB == nil {
			tmpB = headA
		} else {
			tmpB = tmpB.Next
		}
		fmt.Println("tmpB:", tmpB)
		fmt.Printf("\n")
		count++
	}

	return tmpA
}
