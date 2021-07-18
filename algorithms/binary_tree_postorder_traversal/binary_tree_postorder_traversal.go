package binary_tree_postorder_traversal

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
func postorderTraversal(root *TreeNode) []int {
	nodes := make([]int, 0)
	stack := make([]*TreeNode, 0)

	if root != nil {
		stack = append(stack, root)
	}

	for len(stack) != 0 {
		p := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if p.Left != nil {
			stack = append(stack, p.Left)
		}

		if p.Right != nil {
			stack = append(stack, p.Right)
		}

		nodes = append(nodes, p.Val)
	}

	return reverse(nodes)
}

func reverse(nodes []int) []int {

	i := 0
	j := len(nodes) - 1
	for i < j {
		nodes[i], nodes[j] = nodes[j], nodes[i]
		i++
		j--
	}

	return nodes
}
