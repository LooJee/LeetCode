package lcof_level_order_tree_iii

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

func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	treeList := []*TreeNode{root}
	level := 1
	size := 1

	for len(treeList) > 0 {
		nodes := treeList[:size]
		tmpSize := size
		var arr []int

		size = 0
		for _, node := range nodes {
			if node.Left != nil {
				treeList = append(treeList, node.Left)
				size++
			}
			if node.Right != nil {
				treeList = append(treeList, node.Right)
				size++
			}

			arr = append(arr, node.Val)
		}

		if level%2 == 0 {
			e := len(arr) - 1
			for i := 0; i < len(arr)/2; i++ {
				arr[i], arr[e-i] = arr[e-i], arr[i]
			}
		}

		treeList = treeList[tmpSize:]
		level++
		res = append(res, arr)
	}

	return res
}
