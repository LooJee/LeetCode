package populatingnextrightpointersineachnode

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	nodes := []*Node{root}

	for len(nodes) > 0 {
		nodeLen := len(nodes)
		for i := 0; i < nodeLen; i++ {
            if i < nodeLen-1 {
                nodes[i].Next = nodes[i+1]
            }
			if nodes[i].Left != nil {
				nodes = append(nodes, nodes[i].Left)
			}

			if nodes[i].Right != nil {
				nodes = append(nodes, nodes[i].Right)
			}
		}

		nodes = nodes[nodeLen:]
	}

	return root
}
