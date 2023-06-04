package serializeanddeserializebinarytree

import (
	"strconv"
	"strings"
)

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

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}

	return doSerialize(root)
}

func doSerialize(root *TreeNode) string {
	if root == nil {
		return "#"
	}

	builder := strings.Builder{}
	builder.WriteString(strconv.FormatInt(int64(root.Val), 10))
	builder.WriteByte(',')

	builder.WriteString(doSerialize(root.Left))
	builder.WriteByte(',')
	builder.WriteString(doSerialize(root.Right))

	return builder.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}

	var idx = 0
	return doDeserialize(strings.Split(data, ","), &idx)
}

func doDeserialize(data []string, idx *int) *TreeNode {
	if len(data) <= *idx || data[*idx] == "#" {
		return nil
	}

	root := &TreeNode{}
	root.Val, _ = strconv.Atoi(data[*idx])

	*idx++
	root.Left = doDeserialize(data, idx)
	*idx++
	root.Right = doDeserialize(data, idx)

	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
