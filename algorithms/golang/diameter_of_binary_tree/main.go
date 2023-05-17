package diameterofbinarytree

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

var maxDiameter int

func diameterOfBinaryTree(root *TreeNode) int {
	maxDiameter = 0
	depth(root)
	return maxDiameter
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := depth(root.Left)
	rightDepth := depth(root.Right)

	maxDiameter = int(math.Max(float64(maxDiameter), float64(leftDepth+rightDepth)))

	return int(math.Max(float64(leftDepth), float64(rightDepth))) + 1
}
