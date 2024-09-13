package minAbsDifInBFS

import (
	"leetcode/node"
	"math"
)

func GetMinimumDifference(root *node.TreeNode) int {
	minDiff := math.MaxInt
	var prev *node.TreeNode // nil

	var inOrderTraverse func(node *node.TreeNode)
	inOrderTraverse = func(node *node.TreeNode) {
		if node == nil {
			return
		}

		inOrderTraverse(node.Left)

		if prev != nil {
			diff := node.Val - prev.Val
			if diff < minDiff {
				minDiff = diff
			}
		}
		prev = node

		inOrderTraverse(node.Right)
	}

	inOrderTraverse(root)
	return minDiff
}
