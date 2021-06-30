package balanced_binary_tree

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
func isBalanced(root *TreeNode) bool {
	return height(root) >= 0
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}

	rightHeight := height(root.Right)
	leftHeight := height(root.Left)

	if rightHeight == -1 || leftHeight == -1 {
		return -1
	}

	if (rightHeight-leftHeight > 1) || (rightHeight-leftHeight < -1) {
		return -1
	}

	if rightHeight > leftHeight {
		return rightHeight + 1
	} else {
		return leftHeight + 1
	}
}
