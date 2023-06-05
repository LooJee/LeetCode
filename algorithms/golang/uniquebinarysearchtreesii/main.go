package uniquebinarysearchtreesii

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

func generateTrees(n int) []*TreeNode {
	return generate(1, n)
}

func generate(left, right int) []*TreeNode {
	if left > right {
		return []*TreeNode{nil}
	}

	roots := make([]*TreeNode, 0)
	for i := left; i <= right; i++ {
		leftNodes := generate(left, i-1)
		rightNodes := generate(i+1, right)

		for _, leftNode := range leftNodes {
			for _, rightNode := range rightNodes {
				roots = append(roots, &TreeNode{
					Val:   i,
					Left:  leftNode,
					Right: rightNode,
				})
			}
		}
	}

	return roots
}
