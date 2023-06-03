package constructbinarytreefrompreorderandinordertraversal

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

func buildTree(preorder []int, inorder []int) *TreeNode {
	return build(preorder, 0, len(preorder)-1, inorder, 0, len(inorder)-1)
}

func build(preorder []int, preorderStart, preorderEnd int, inorder []int, inorderStart, inorderEnd int) *TreeNode {
	if inorderStart > inorderEnd {
		return nil
	}

	var (
		root           = preorder[preorderStart]
		rootInorderIdx = -1
	)

	// 找到中序遍历的 root 索引
	for i := inorderStart; i <= inorderEnd; i++ {
		if inorder[i] == root {
			rootInorderIdx = i
		}
	}

	// 中序遍历的左子树区间为 leftInorder = [inorderStart, rootInorderIdx-1]
	// 中序遍历的右子树区间为 rightInorder = [rootInorderIdx+1, inorderEnd]
	// 前序遍历的左子树区间为 [preorderStart+1, preorderStart+len(leftInorder)]
	// 前序遍历的右子树区间为 [preorderStart+len(leftInorder)+1, preorderEnd]
	return &TreeNode{
		Val:   root,
		Left:  build(preorder, preorderStart+1, preorderStart+(rootInorderIdx-inorderStart), inorder, inorderStart, rootInorderIdx-1),
		Right: build(preorder, preorderEnd-(inorderEnd-rootInorderIdx-1), preorderEnd, inorder, rootInorderIdx+1, inorderEnd),
	}
}
