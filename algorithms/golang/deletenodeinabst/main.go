package deletenodeinabst

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

func deleteNode(root *TreeNode, key int) *TreeNode {
	return doDelete(root, key)
}

func doDelete(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == key {
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		} else {
			// 如果左右子树都存在，则找到右子树的最小值替换当前节点的位置
			minNode := findMin(root.Right)
			// 将右子树中原来的 minNode 删除
			minNode.Right = doDelete(root.Right, minNode.Val)
			minNode.Left = root.Left
			root = minNode
		}
	} else if root.Val < key {
		root.Right = doDelete(root.Right, key)
	} else if root.Val > key {
		root.Left = doDelete(root.Left, key)
	}

	return root
}

func findMin(root *TreeNode) *TreeNode {
	if root.Left == nil {
		return root
	}

	return findMin(root.Left)
}
