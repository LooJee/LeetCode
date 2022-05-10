package lcof_copy_complex_list

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomListByMap(head *Node) *Node {
	if head == nil {
		return nil
	}

	m := make(map[*Node]*Node)
	dummyHead := &Node{}

	tmp := dummyHead
	for t := head; t != nil; t = t.Next {
		node, ok := m[t]
		if ok {
			tmp.Next = node
		} else {
			node = &Node{Val: t.Val}
			tmp.Next = node
			m[t] = node
		}

		if t.Random != nil {
			if random, ok := m[t.Random]; ok {
				node.Random = random
			} else {
				random = &Node{Val: t.Random.Val}
				node.Random = random
				m[t.Random] = random
			}
		}

		tmp = tmp.Next
	}

	return dummyHead.Next
}

func copyRandomListBySibling(head *Node) *Node {
	if head == nil {
		return nil
	}

	for t := head; t != nil; t = t.Next.Next {
		t.Next = &Node{
			Val:  t.Val,
			Next: t.Next,
		}
	}

	for t := head; t != nil; t = t.Next.Next {
		if t.Random != nil {
			t.Next.Random = t.Random.Next
		}
	}

	dummyHead := &Node{}
	tmp := dummyHead
	for t := head; t != nil; t = t.Next {
		tmp.Next = t.Next
		t.Next = t.Next.Next
		tmp = tmp.Next
	}

	return dummyHead.Next
}
