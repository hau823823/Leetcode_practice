package symmetricTree

import "leetcode/node"

func IsSymmetric(root *node.TreeNode) bool {
	queue := []*node.TreeNode{root.Left, root.Right}

	for len(queue) != 0 {
		leftNode := queue[0]
		rightNode := queue[1]
		queue = queue[2:]

		if leftNode != nil && rightNode != nil {
			if leftNode.Val == rightNode.Val {
				queue = append(queue, leftNode.Left, rightNode.Right, leftNode.Right, rightNode.Left)
			} else {
				return false
			}
		} else if leftNode != nil && rightNode == nil {
			return false
		} else if leftNode == nil && rightNode != nil {
			return false
		}
	}

	return true
}

func IsSymmetricRecursive(root *node.TreeNode) bool {
	var isMirror func(*node.TreeNode, *node.TreeNode) bool
	isMirror = func(left, right *node.TreeNode) bool {
		if left == nil && right == nil {
			return true
		}
		if left == nil || right == nil || left.Val != right.Val {
			return false
		}
		return isMirror(left.Left, right.Right) && isMirror(left.Right, right.Left)
	}

	return isMirror(root.Left, root.Right)
}