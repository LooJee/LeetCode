package validatebinarysearchtree

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

// 左子树的所有值都必须小于root，右子树的所有值都必须大于root
func isValidBST(root *TreeNode) bool {
	return traverse(root, nil, nil)
}

func traverse(root *TreeNode, min, max *TreeNode) bool {
	if root == nil {
		return true
	}

	if min != nil && min.Val >= root.Val {
		return false
	}

	if max != nil && max.Val <= root.Val {
		return false
	}

	return traverse(root.Left, min, root) && traverse(root.Right, root, max)
}
