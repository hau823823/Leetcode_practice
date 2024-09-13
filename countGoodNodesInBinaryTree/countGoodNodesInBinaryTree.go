package countGoodNodesInBinaryTree

import "leetcode/node"

func GoodNodes(root *node.TreeNode) int {

	var dfs func(*node.TreeNode, int) int
	dfs = func(t *node.TreeNode, parent int) int {
		if t == nil {
			return 0
		}

		res := 1
		max := t.Val

		if parent > t.Val {
			res = 0
			max = parent
		}

		res += dfs(t.Left, max)
		res += dfs(t.Right, max)

		return res
	}

	return dfs(root, root.Val)
}
