package convert_sorted_array_to_binary_search_tree

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
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	tree := &TreeNode{}
	mid := len(nums) / 2
	tree.Val = nums[mid]
	tree.Left = sortedArrayToBST(nums[:mid])
	tree.Right = sortedArrayToBST(nums[mid+1:])

	return tree
}
