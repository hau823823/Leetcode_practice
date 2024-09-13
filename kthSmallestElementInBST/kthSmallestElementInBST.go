package kthSmallestElementInBST

import "leetcode/node"

// inorder recursive
func KthSmallest(root *node.TreeNode, k int) int {
    var inorder func(*node.TreeNode) []int
	inorder = func(root *node.TreeNode) []int {
		if root == nil {
			return []int{}
		}
		res := inorder(root.Left)
		res = append(res, root.Val)
		res = append(res, inorder(root.Right)...)

		return res
	}
	tmp := inorder(root)
	return tmp[k-1]
}

// solution2
func KthSmallest2(root *node.TreeNode, k int) int {
	count := k
	var res int
	var dfs func(*node.TreeNode)
	dfs = func(root *node.TreeNode) {
		if root == nil {
			return 
		}
		dfs(root.Left)
		count --
		if count == 0 {
			res = root.Val
		}
		dfs(root.Right)
	}
	dfs(root)
	return  res
}

// count
func KthSmallestCount(root *node.TreeNode, k int) int {
	var helper func(*node.TreeNode, int) int
	helper = func(root *node.TreeNode, k int) int {
		cnt := countFunc(root.Left)
		if k <= cnt {
			return helper(root.Left, k)
		} else if k > cnt + 1 {
			return helper(root.Right, k - cnt - 1)
		}
		return root.Val
	}

	return helper(root, k)
}

func countFunc(root *node.TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + countFunc(root.Left) + countFunc(root.Right)
}

// follow up: the BST is modified often 
// (i.e., we can do insert and delete operations) and you need to find the kth smallest frequently)
type NewTreeNode struct {
	Val   int
	Count int
	Left  *NewTreeNode
	Right *NewTreeNode
}

func Build(root *node.TreeNode) *NewTreeNode {
	if root == nil {
		return nil
	}

	node := &NewTreeNode{Val: root.Val, Count: 1}
	node.Left = Build(root.Left)
	node.Right = Build(root.Right)

	if node.Left != nil {
		node.Count += node.Left.Count
	}
	if node.Right != nil {
		node.Count += node.Right.Count
	}

	return node
}

func Helper(node *NewTreeNode, k int) int {
	if node.Left != nil {
		cnt := node.Left.Count
		if k <= cnt {
			return Helper(node.Left, k)
		} else if k > cnt + 1 {
			return Helper(node.Right, k - cnt -1)
		}
		return node.Val
	} else {
		if k == 1 {
			return node.Val
		}
		return Helper(node.Right, k - 1)
	}
}

func KthSmallestOP(root *node.TreeNode, k int) int {
	node := Build(root)
	return Helper(node, k)
}