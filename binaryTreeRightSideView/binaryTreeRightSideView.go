package binaryTreeRightSideView

import "leetcode/node"

func RightSideView(root *node.TreeNode) []int {
	if root == nil {
		return []int{}
	}

	queue := []*node.TreeNode{root}
	result := []int{}
	for len(queue) > 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		result = append(result, queue[l - 1].Val)
		queue = queue[l:]
	}

	return result
}
