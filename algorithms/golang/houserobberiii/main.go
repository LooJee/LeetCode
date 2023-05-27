package houserobberiii

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

func rob(root *TreeNode) int {
	dp_1, _ := doRob(root)

	return dp_1
}

func doRob(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}

	var (
		dp_l_1, dp_l_2 = 0, 0
		dp_r_1, dp_r_2 = 0, 0
	)

    dp_l_1, dp_l_2 = doRob(root.Left)
    dp_r_1, dp_r_2 = doRob(root.Right)

	return int(math.Max(float64(dp_l_1+dp_r_1), float64(dp_l_2+dp_r_2+root.Val))), dp_l_1 + dp_r_1
}
