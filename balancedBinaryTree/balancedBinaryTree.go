package balancedBinaryTree

import (
	"leetcode/node"
	"math"
)

func IsBalanced(root *node.TreeNode) bool {
	gap := 0.0

	var dfs func(*node.TreeNode) (int, bool)
	dfs = func(t *node.TreeNode) (int, bool) {
		if t == nil{
			return 0, true
		}

		left, isLeftBalanced := dfs(t.Left)
		right, isRightBalanced := dfs(t.Right)
		gap = math.Abs(float64(left - right))

		if isLeftBalanced && isRightBalanced && gap <= 1 {
			return max(left, right) + 1, true
		}
		return -1, false
	}

	_, res := dfs(root)
	return res
}

func IsBalancedOp(root *node.TreeNode) bool {
	gap  := 0.0
    ans := true

    var dfs func(*node.TreeNode)int
    dfs = func(node *node.TreeNode)int{
        if node == nil || !ans{
            return 0
        }

        left := dfs(node.Left)
        right := dfs(node.Right)
		gap = math.Abs(float64(left - right))

        if gap >  1{
            ans = false
        }
        return max(left,right) +1
    }

    dfs(root)
    return ans
    
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
