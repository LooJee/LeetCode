package lcof_max_depth

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	return down(root)
}

func down(node *TreeNode) int {
	if node == nil {
		return 0
	}

	return max(down(node.Left)+1, down(node.Right)+1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
