package binaryTreeZigzagLevelOrderTraversal

import "leetcode/node"

func ZigzagLevelOrder(root *node.TreeNode) [][]int {
	result := [][]int{}

	var bfs func(*node.TreeNode, int)
	bfs = func(root *node.TreeNode, level int) {
		if root == nil {
			return
		}
		if level >= len(result) {
			result = append(result, []int{})
		}
		if level % 2 == 0 {
			result[level] = append(result[level], root.Val)
		} else {
			result[level] = append([]int{root.Val}, result[level]...)
		}
		bfs(root.Left, level+1)
		bfs(root.Right, level+1)

	}
	bfs(root, 0)

    return result
}

func ZigZagLevelOrderLoop(root *node.TreeNode) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}

	queue := []*node.TreeNode{root}
	level := 0
	for len(queue) > 0 {
		l := len(queue)
		tmp := []int{}
		for i := 0; i < l; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
			if level % 2 == 0 {
				tmp = append(tmp, queue[i].Val)
			} else {
				tmp = append([]int{queue[i].Val}, tmp...)
			}
		}
		queue = queue[l:]
		result = append(result, tmp)
		level ++
	}

	return result
}