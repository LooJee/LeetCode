package lcof_is_balanced

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

var balanced = true

func isBalanced(root *TreeNode) bool {
	balanced = true
	down(root)
	return balanced
}

func down(node *TreeNode) int {
	if node == nil {
		return 0
	}

	lDeep := down(node.Left)
	rDeep := down(node.Right)

	if Abs(lDeep-rDeep) > 1 {
		balanced = false
		return -1
	}

	return Max(lDeep, rDeep) + 1
}

func Abs(c int) int {
	if c < 0 {
		return -c
	} else {
		return c
	}
}

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
