package maximumDepthBT

import (
	"fmt"
	"leetcode/node"
)

func MaxDepth(root *node.TreeNode) int {
	if root == nil {
		return 0
	}
	maxLeft := MaxDepth(root.Left)
	maxRight := MaxDepth(root.Right)

	res := 0
	if maxLeft > maxRight {
		res += maxLeft
	} else {
		res += maxRight
	}

	return res + 1
}

// beautify
func MaxDepthOp(root *node.TreeNode) int {
	var maxDepth func(*node.TreeNode) int

	maxDepth = func(root *node.TreeNode) int {
		if root == nil {
			return 0
		}
		maxLeft := maxDepth(root.Left)
		maxRight := maxDepth(root.Right)

		res := maxRight
		if maxLeft > maxRight {
			res = maxLeft
		}

		return res + 1
	}

	return maxDepth(root)
}

func MaxDepthTEST(root *node.TreeNode, count int) int {
	if root == nil {
		return 0
	}
	count++
	_ = MaxDepthTEST(root.Left, 0)
	_ = MaxDepthTEST(root.Right, 0)
	
	fmt.Println(count)

	return 0
}
