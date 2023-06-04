package constructbinarytreefrominorderandpostordertraversal

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

func buildTree(inorder []int, postorder []int) *TreeNode {
	return build(inorder, 0, len(inorder)-1, postorder, 0, len(postorder)-1)
}

func build(inorder []int, inorderStart, inorderEnd int, postorder []int, postorderStart, postorderEnd int) *TreeNode {
	if postorderStart > postorderEnd {
		return nil
	}

	var (
		root    = postorder[postorderEnd]
		rootIdx = -1
	)

	// 在中序中找到 root， 分割左右子树
	for i := inorderStart; i <= inorderEnd; i++ {
		if inorder[i] == root {
			rootIdx = i
			break
		}
	}

	// 中序左子树： leftInorder = [inorderStart, rootIdx-1]
	// 中序右子树: rightInorder = [rootIdx+1, inorderEnd]
	// 后序左子树: leftPostorder = [postorderStart, postorderStart+(rootIdx-inorderStart-1)]
	// 后序右子树: rightInorder = [postorderEnd-(inorderEnd-rootIdx), postorderEnd-1]

	return &TreeNode{
		Val:   root,
		Left:  build(inorder, inorderStart, rootIdx-1, postorder, postorderStart, postorderStart+(rootIdx-inorderStart-1)),
		Right: build(inorder, rootIdx+1, inorderEnd, postorder, postorderEnd-(inorderEnd-rootIdx), postorderEnd-1),
	}
}
