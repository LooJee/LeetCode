package countcompletetreenodes

import "math"

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

// 可以先判断一下完全二叉树是否是满二叉树，如果是满二叉树可以使用公式 2^n - 1 来计算点数量
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var (
		leftTree    = root
		rightTree   = root
		leftHeight  = 0
		rightHeight = 0
	)

	// 分别遍历左枝和右枝，如果相同，则是满二叉树
	for leftTree != nil {
		leftTree = leftTree.Left
		leftHeight++
	}

	for rightTree != nil {
		rightTree = rightTree.Right
		rightHeight++
	}

	if leftHeight == rightHeight {
		return int(math.Pow(2, float64(leftHeight))) - 1
	}

	// 如果不是满二叉树，则逐个遍历
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}
