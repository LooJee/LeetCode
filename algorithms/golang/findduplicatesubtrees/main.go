package findduplicatesubtrees

import "fmt"

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

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	var (
		memos = make(map[string]int)
		nodes = make([]*TreeNode, 0)
	)

	find(root, memos, &nodes)

	return nodes
}

func find(root *TreeNode, memos map[string]int, nodes *[]*TreeNode) string {
	if root == nil {
		return ""
	}

	left := find(root.Left, memos, nodes)
	right := find(root.Right, memos, nodes)

	subTree := fmt.Sprintf("%d,%s,%s", root.Val, left, right)
	memos[subTree]++

	if memos[subTree] == 2 {
		*nodes = append(*nodes, root)
	}

	return subTree
}
