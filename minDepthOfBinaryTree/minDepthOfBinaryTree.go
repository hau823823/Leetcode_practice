package minDepthOfBinaryTree

import "leetcode/node"

func MinDepth(root *node.TreeNode) int {

	var depth func(*node.TreeNode) int
	depth = func(node *node.TreeNode) int {
		if node == nil {
			return 0
		}
		minLeft := depth(node.Left)
		minRight := depth(node.Right)

		if minLeft == 0 || minRight == 0 {
			return minLeft + minRight + 1
		}
		return min(minLeft, minRight) + 1
	}

	return depth(root)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}