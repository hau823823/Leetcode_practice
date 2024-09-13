package constructBinaryTreeFromPreOrderAndInOrder

import (
	"leetcode/node"
)

// preorder: ROOT -> LEFT -> RIGHT
// inorder:  LEFT -> ROOT -> RIGHT
func BuildTree(preorder []int, inorder []int) *node.TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	// 查找 root 在 inorder 數列中的位置
	index := 0
	for i := 0; i < len(inorder); i++ {
		if preorder[0] == inorder[i] {
			index = i
			break
		}
	}

	result := &node.TreeNode{
		Val:   preorder[0],
		Left:  BuildTree(preorder[1:1+index], inorder[:index]),
		Right: BuildTree(preorder[1+index:], inorder[1+index:]),
	}

	return result
}

// use map optimize
func BuildTreeOP(preorder []int, inorder []int) *node.TreeNode {
	inMap := make(map[int]int, len(inorder))
	for i, val := range inorder {
		inMap[val] = i
	}

	var build func([]int, map[int]int, int, int, int) *node.TreeNode
	build = func(preorder []int, inMap map[int]int, l, r, preIdx int) *node.TreeNode {
		if l > r {
			return nil
		}

		val := preorder[preIdx]
		root := &node.TreeNode{Val: val, Left: nil, Right: nil}
		if l == r {
			return root
		}

		idx := inMap[val]

		root.Left = build(preorder, inMap, l, idx-1, preIdx+1)
		root.Right = build(preorder, inMap, idx+1, r, preIdx+idx-l+1)
		return root

	}

	return build(preorder, inMap, 0, len(inorder)-1, 0)
}
