package tree

/*
   this function deletes node from the tree, ensuring that
   the deleted node is replaced by the right most node in the
   tree and that the tree shrinks in height in that process
*/
func RightMostDeep(n, m *Node) {
	var temp *Node
	var queue = make([]*Node, 0)
	queue = append(queue, n)
	for len(queue) > 0 {
		temp = queue[len(queue)-1]
		if temp.Left == m {
			temp.Left = nil
			return
		} else {
			queue = append(queue, temp.Left)
		}
		if temp.Right == m {
			temp.Right = nil
			return
		} else {
			queue = append(queue, temp.Right)
		}
	}
	return
}
func DeleteNode(root, key *Node) {
	var queue = make([]*Node, 0)
	var t *Node
	var temp *Node
	queue = append(queue, root)
	for len(queue) > 0 {
		// itrating throught the queue to find 1. the right most and
		// deepest node, 2. the node which we wish to replace.
		temp = queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		if key.Val == temp.Val {
			t = temp
		}
		if temp.Right != nil {
			queue = append(queue, temp.Right)
		}
		if temp.Left != nil {
			queue = append(queue, temp.Left)
		}
	}
	t.Val = temp.Val
	RightMostDeep(root, temp) // deleting the right most and deepest node
	return
}
