package invertBinaryTree

import "leetcode/node"

func InvertTree(root *node.TreeNode) *node.TreeNode {
	if root == nil {
		return root
	}
	left := InvertTree(root.Left)
	root.Left = InvertTree(root.Right)
	root.Right = left

	return root
}
