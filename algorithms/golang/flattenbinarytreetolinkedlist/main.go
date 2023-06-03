package flattenbinarytreetolinkedlist

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

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	flatten(root.Left)
	flatten(root.Right)

	if root.Left == nil {
		return
	}

	front := root.Left

	for front.Right != nil {
		front = front.Right
	}

	front.Right = root.Right
	root.Right = root.Left
	root.Left = nil
}
