package binarysearchtreetogreatersumtree

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

// 从右子树开始遍历，
func bstToGst(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	var sum = 0

	traverse(root, &sum)

	return root
}

func traverse(root *TreeNode, sum *int) {
	if root == nil {
		return
	}

	traverse(root.Right, sum)
	root.Val += *sum
	*sum = root.Val
	traverse(root.Left, sum)
}
