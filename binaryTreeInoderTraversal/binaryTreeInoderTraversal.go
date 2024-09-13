package binaryTreeInoderTraversal

import "leetcode/node"

// recursive
func InorderTraversalRe(root *node.TreeNode) []int {

    var inorder func(*node.TreeNode) []int
    
    inorder = func(root *node.TreeNode) []int {

        if root == nil {
            return []int{}
        }

        //if root.Left == nil && root.Right == nil {
        //    return []int{root.Val}
        //}

        res := inorder(root.Left)
        res = append(res, root.Val)
        res = append(res, inorder(root.Right)...)

        return  res
    }

    return inorder(root)
}

// iterative
func InorderTraversalIr(root *node.TreeNode) []int {
    res := []int{}
    stack := []*node.TreeNode{}
    cur := root 

    for cur != nil || len(stack) != 0 {
        for cur != nil {
            stack = append(stack, cur)
            cur = cur.Left
        }

        cur = stack[len(stack) - 1]
        stack = stack[:len(stack) - 1]
        res = append(res, cur.Val)
        cur = cur.Right
    }
    
    return res
}