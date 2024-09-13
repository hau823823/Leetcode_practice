package populatingNextRightPtrInEachNode

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func Connect(root *Node) *Node {
	if root == nil {
		return root
	}

	queue := []*Node{root}
	for len(queue) > 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
			if i == l - 1 {
				queue[i].Next = nil
			} else {
				queue[i].Next = queue[i+1]
			}
		}
		queue = queue[l:]
	}

	return root
}