package subTreeOfAnotherTree

import "leetcode/node"

func IsSubtree(root *node.TreeNode, subRoot *node.TreeNode) bool {

	if root == nil {
		return false
	}

	var dfs func(*node.TreeNode, *node.TreeNode) bool
	dfs = func(root *node.TreeNode, subRoot *node.TreeNode) bool {
		if root == nil && subRoot == nil {
			return true
		}
		if root == nil || subRoot == nil || root.Val != subRoot.Val {
			return false
		}
		return dfs(root.Left, subRoot.Left) && dfs(root.Right, subRoot.Right)
	}

	return dfs(root, subRoot) || IsSubtree(root.Left, subRoot) || IsSubtree(root.Right, subRoot)
}
