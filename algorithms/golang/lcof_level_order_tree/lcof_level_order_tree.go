package lcof_level_order_tree

import "container/list"

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

func levelOrderByList(root *TreeNode) []int {
	var arr []int

	if root == nil {
		return arr
	}
	treeList := list.New()
	treeList.PushBack(root)

	for treeList.Len() > 0 {
		node := treeList.Remove(treeList.Front()).(*TreeNode)
		if node.Left != nil {
			treeList.PushBack(node.Left)
		}
		if node.Right != nil {
			treeList.PushBack(node.Right)
		}

		arr = append(arr, node.Val)
	}

	return arr
}

func levelOrderBySlice(root *TreeNode) []int {
	var arr []int

	if root == nil {
		return arr
	}
	treeList := make([]*TreeNode, 0)
	treeList = append(treeList, root)

	for len(treeList) > 0 {
		node := treeList[0]
		if node.Left != nil {
			treeList = append(treeList, node.Left)
		}
		if node.Right != nil {
			treeList = append(treeList, node.Right)
		}
		treeList = treeList[1:]

		arr = append(arr, node.Val)
	}

	return arr
}
