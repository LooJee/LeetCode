package symmetric_tree

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
func isSymmetric(root *TreeNode) bool {
	return internalIsSymmetric(root.Left, root.Right)
}

func internalIsSymmetric(left, right *TreeNode) bool {
	if left == nil && right != nil || left != nil && right == nil {
		return false
	}

	if left == nil && right == nil {
		return true
	}

	return left.Val == right.Val && internalIsSymmetric(left.Left, right.Right) && internalIsSymmetric(left.Right, right.Left)
}
