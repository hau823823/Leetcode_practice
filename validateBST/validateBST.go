package validateBST

import (
	"leetcode/node"
	"math"
)

// recursive
func IsValidBST(root *node.TreeNode) bool {
	var isValid func(*node.TreeNode, int, int) bool

	isValid = func(root *node.TreeNode, min int, max int) bool {
		if root == nil {
			return true
		}
		if root.Val <= min || root.Val >= max {
			return false
		}
		
		return isValid(root.Left, min, root.Val) && isValid(root.Right, root.Val, max)
	}

	return isValid(root, math.MinInt, math.MaxInt)
}
