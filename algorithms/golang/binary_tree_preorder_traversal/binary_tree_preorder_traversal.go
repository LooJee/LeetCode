package binary_tree_preorder_traversal

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
func preorderTraversal(root *TreeNode) []int {
	nodes := make([]int, 0)
	if root == nil {
		return nodes
	}

	stack := make([]*TreeNode, 0)
	stack = append(stack, root)

	for len(stack) != 0 {
		node := stack[len(stack)-1]
		nodes = append(nodes, node.Val)
		stack = stack[:len(stack)-1]
		if node.Right != nil {
			stack = append(stack, node.Right)
		}

		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}

	return nodes
}
