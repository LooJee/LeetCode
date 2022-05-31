package lcof_path_sum

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

func pathSum(root *TreeNode, target int) [][]int {
	result := make([][]int, 0)

	path(root, &result, 0, target, []int{})

	return result
}

func path(node *TreeNode, result *[][]int, sum, target int, vals []int) {
	if node == nil {
		return
	}

	sum += node.Val
	vals = append(vals, node.Val)

	if node.Left == nil && node.Right == nil && sum == target {
		*result = append(*result, append([]int{}, vals...))
		return
	}

	path(node.Left, result, sum, target, vals)
	path(node.Right, result, sum, target, vals)
}
