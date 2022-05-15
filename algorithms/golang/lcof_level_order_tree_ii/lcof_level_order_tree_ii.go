package lcof_level_order_tree_ii

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

	treeList := make([]*TreeNode, 0)
	treeList = append(treeList, root)
	tmpSize, size := 1, 1

	for len(treeList) > 0 {
		tmpSize = size
		tmp := treeList[:size]
		arr := make([]int, 0)
		size = 0
		for _, node := range tmp {
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
		res = append(res, arr)
		treeList = treeList[tmpSize:]
	}

	return res
}
