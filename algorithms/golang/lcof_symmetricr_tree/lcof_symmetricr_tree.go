package lcof_symmetricr_tree

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

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return check(root.Left, root.Right)
}

func check(l, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}

	if l == nil && r != nil || l != nil && r == nil {
		return false
	}

	if l.Val != r.Val {
		return false
	}
	return check(l.Left, r.Right) && check(l.Right, r.Left)
}
