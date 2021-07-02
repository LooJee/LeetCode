package minimum_depth_of_binary_tree

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	minDep := math.MaxInt32
	if root.Left != nil {
		minDep = min(minDepth(root.Left), minDep)
	}

	if root.Right != nil {
		minDep = min(minDepth(root.Right), minDep)
	}

	return minDep + 1
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
