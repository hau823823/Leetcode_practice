package diameterOfBinaryTree

import "leetcode/node"

func DiameterOfBinaryTree(root *node.TreeNode) int {
	maxLength := 0

	var dfs func(*node.TreeNode) int
	dfs = func(t *node.TreeNode) int {
		if t == nil {
			return 0
		}

		left := dfs(t.Left)
		right := dfs(t.Right)
		maxLength = max(maxLength, left + right)

		return max(left, right) + 1
	}

	dfs(root)
	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}