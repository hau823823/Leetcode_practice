package lowestCommonAncestorOfBST

import "leetcode/node"

// recursive
func LowestCommonAncestorRecursive(root, p, q *node.TreeNode) *node.TreeNode {
	if p.Val < root.Val && q.Val < root.Val {
		return LowestCommonAncestorRecursive(root.Left, p, q)
	}	
	if p.Val > root.Val && q.Val > root.Val {
		return LowestCommonAncestorRecursive(root.Right, p, q)
	}

	return root
}

// iterative
func LowestCommonAncestorIterative(root, p, q *node.TreeNode) *node.TreeNode {
	for root != nil {
		if p.Val < root.Val && q.Val < root.Val {
			root = root.Left
		} else if p.Val > root.Val && q.Val > root.Val {
			root = root.Right
		} else {
			return root
		}
	}

	return nil
}