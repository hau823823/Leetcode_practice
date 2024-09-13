package convertArray2BST

import "leetcode/node"

func SortedArrayToBST(nums []int) *node.TreeNode {
    var traversal func([]int, int, int) *node.TreeNode

    traversal = func(nums []int, left, right int) *node.TreeNode {
        if left > right {
            return nil
        }

        mid := (left + right) / 2
        root := &node.TreeNode{Val: nums[mid]}
        root.Left = traversal(nums, left, mid-1)
        root.Right = traversal(nums, mid+1, right)

        return root
    }

    root := traversal(nums, 0, len(nums)-1)
    return root
}