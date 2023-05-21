package minimum_depth_of_binary_tree

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
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}
	depth := 1

	for len(queue) > 0 {
		queLen := len(queue)

		for i := 0; i < queLen; i++ {
			node := queue[0]
			queue = queue[1:]

			if node.Left == nil && node.Right == nil {
				return depth
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		depth++
	}

	return depth
}
