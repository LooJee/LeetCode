package lcof_kth_largest

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

func kthLargest(root *TreeNode, k int) int {
	var val int
	dfs(root, &k, &val)

	return val
}

func dfs(root *TreeNode, k, val *int) {
	if root == nil {
		return
	}

	dfs(root.Right, k, val)
	if *k > 0 {
		*k--
	} else {
		return
	}

	if *k == 0 {
		*val = root.Val
		return
	}

	dfs(root.Left, k, val)
}
