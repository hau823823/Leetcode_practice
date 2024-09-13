package node

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var NULL = -1 << 63

// List
func Ints2ListNode(ints []int) *ListNode {
	if len(ints) == 0 {
		return nil
	}

	res := &ListNode{Val: ints[0]}
	tmp := res
	for i := 1; i < len(ints); i++ {
		tmp.Next = &ListNode{Val: ints[i]}
		tmp = tmp.Next
	}

	return res
}

func ListNode2Ints(list *ListNode) []int {
	res := []int{}

	for list != nil {
		res = append(res, list.Val)
		list = list.Next
	}

	return res
}

// Tree
func Ints2TreeNode(ints []int) *TreeNode {
	n := len(ints)
	if n == 0 {
		return nil
	}

	root := &TreeNode{Val: ints[0]}
	queue := make([]*TreeNode, 1, n*2)
	queue[0] = root

	i := 1
	for i < n {
		node := queue[0]
		queue = queue[1:]
		if i < n && ints[i] != NULL {
			node.Left = &TreeNode{Val: ints[i]}
			queue = append(queue, node.Left)
		}
		i++
		if i < n && ints[i] != NULL {
			node.Right = &TreeNode{Val: ints[i]}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

// pre order
// 中 -> 左 -> 右
func Tree2PreOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	res := []int{root.Val}
	res = append(res, Tree2PreOrder(root.Left)...)
	res = append(res, Tree2PreOrder(root.Right)...)
	return res
}

func Tree2PreOrderStack(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}

	stack := []*TreeNode{root}
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, cur.Val)

		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}
		if cur.Left != nil {
			stack = append(stack, cur.Left)
		}
	}

	return res
}

// in order
// 左 -> 中 -> 右
func Tree2InOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	/**
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	*/
	res := Tree2InOrder(root.Left)
	res = append(res, root.Val)
	res = append(res, Tree2InOrder(root.Right)...)
	return res
}

func Tree2InOrderStack(root *TreeNode) []int {
	res := []int{}
	stack := []*TreeNode{}
	cur := root

	for cur != nil || len(stack) != 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}

		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, cur.Val)
		cur = cur.Right
	}

	return res
}

// post order
// 左 -> 右 -> 中
func Tree2PostOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	res := Tree2PostOrder(root.Left)
	res = append(res, Tree2PostOrder(root.Right)...)
	res = append(res, root.Val)
	return res
}

func Tree2PostOrderStack(root *TreeNode) []int {
	res := []int{}
	stack := []*TreeNode{}
	cur := root
	var lastVisited *TreeNode

	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}

		peekNode := stack[len(stack)-1]
		if peekNode.Right != nil && lastVisited != peekNode.Right {
			cur = peekNode.Right
		} else {
			stack = stack[:len(stack)-1]
			res = append(res, peekNode.Val)
			lastVisited = peekNode
		}
	}

	return res
}
