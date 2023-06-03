package maximumbinarytree

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

func constructMaximumBinaryTree(nums []int) *TreeNode {
	return build(nums, 0, len(nums)-1)
}

func build(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}

	var (
		max    = math.MinInt
		maxIdx = math.MinInt
	)
	for i := left; i <= right; i++ {
		if nums[i] > max {
			max = nums[i]
			maxIdx = i
		}
	}

	return &TreeNode{
		Val:   max,
		Left:  build(nums, left, maxIdx-1),
		Right: build(nums, maxIdx+1, right),
	}
}
