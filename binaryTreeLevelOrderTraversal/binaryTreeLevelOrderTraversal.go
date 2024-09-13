package binaryTreeLevelOrderTraversal

import "leetcode/node"

func LevelOrderRecursive(root *node.TreeNode) [][]int {
	result := [][]int{}
	var bfs func(*node.TreeNode, int)
	bfs = func(root *node.TreeNode, level int) {
		if root == nil {
			return
		}
		if level >= len(result) {
			result = append(result, []int{})
		}
		result[level] = append(result[level], root.Val)
		bfs(root.Left, level+1)
		bfs(root.Right, level+1)
	}
	bfs(root, 0)

    return result
}

func LevelOrderLoop(root * node.TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	queue := []*node.TreeNode{root}
	result := [][]int{}

	for len(queue)  > 0 {
		l := len(queue)
		tmp := make([]int, 0, l)
		for i := 0; i < l; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
			tmp = append(tmp, queue[i].Val)
		}
		queue = queue[l:]
		result = append(result, tmp)
	}

	return result
}