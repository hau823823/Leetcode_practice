package sameTree

import "leetcode/node"

func IsSameTree(p *node.TreeNode, q *node.TreeNode) bool {

	var dfs func(*node.TreeNode, *node.TreeNode) bool
	dfs = func(a *node.TreeNode, b *node.TreeNode) bool {
		if a == nil && b == nil {
			return true
		}

		if a == nil || b == nil || a.Val != b.Val {
			return false
		}

		left := dfs(a.Left, b.Left)
		rght := dfs(a.Right, b.Right)

		return left && rght
	}

	return dfs(p, q)
}