package constructbinarytreefrompreorderandpostordertraversal

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

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	return build(preorder, 0, len(preorder)-1, postorder, 0, len(postorder)-1)
}

func build(preorder []int, preorderStart, preorderEnd int, postorder []int, postorderStart, postorderEnd int) *TreeNode {
	if preorderStart > preorderEnd || postorderStart > postorderEnd {
		return nil
	}

	// preorder 的第一个元素是root， postorder 的最后一个元素是 root
	// preorder 的第二个元素是左子树 root
	// postorder 的倒数第二个元素是右子树 root

	var (
		root                                = preorder[preorderStart]
		preLeftChildIdx, postLeftChildIdx   = preorderStart + 1, -1
		postRightChildIdx, preRightChildIdx = postorderEnd - 1, -1
		leftChild, rightChild               *TreeNode
		preLeftEnd                          = preorderEnd
		postRightStart                      = postorderStart
	)

	// 如果前序左子树存在，则找到后序的左子树序号
	if preLeftChildIdx <= preorderEnd {
		for i := postorderStart; i <= postorderEnd; i++ {
			if postorder[i] == preorder[preLeftChildIdx] {
				postLeftChildIdx = i
				postRightStart = i + 1
				break
			}
		}
	}

	// 如果后序右子树存在，则找到前序的右子树序号
	if postRightChildIdx >= postorderStart && postorder[postRightChildIdx] != preorder[preLeftChildIdx] {
		for i := preorderStart; i <= preorderEnd; i++ {
			if preorder[i] == postorder[postRightChildIdx] {
				preRightChildIdx = i
				preLeftEnd = i - 1
				break
			}
		}
	}

	if preLeftChildIdx <= preorderEnd {
		leftChild = build(preorder, preLeftChildIdx, preLeftEnd, postorder, postorderStart, postLeftChildIdx)
	}

	if postRightChildIdx >= postorderStart {
		rightChild = build(preorder, preRightChildIdx, preorderEnd, postorder, postRightStart, postRightChildIdx)
	}

	return &TreeNode{
		Val:   root,
		Left:  leftChild,
		Right: rightChild,
	}
}
