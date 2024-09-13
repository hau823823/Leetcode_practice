package copyListWithRandomPointer

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func CopyRandomList(head *Node) *Node {
	old2NeweMap := make(map[*Node]*Node)

	ptr := head
	for ptr != nil {
		clone := &Node{Val: ptr.Val}
		old2NeweMap[ptr] = clone
		ptr = ptr.Next
	}

	ptr = head
	for ptr != nil {
		clone := old2NeweMap[ptr]
		clone.Next = old2NeweMap[ptr.Next]
		clone.Random = old2NeweMap[ptr.Random]
		ptr = ptr.Next
	}
	
	return old2NeweMap[head]
}
