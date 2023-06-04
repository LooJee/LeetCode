package kthsmallestelementinabst

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

func kthSmallest(root *TreeNode, k int) int {
	var smallest int

	traverse(root, &k, &smallest)

	return smallest
}

func traverse(root *TreeNode, k *int, smallest *int) {
	if root == nil {
		return
	}

	traverse(root.Left, k, smallest)
	*k--
	if *k == 0 {
		*smallest = root.Val
	}
	traverse(root.Right, k, smallest)
}
