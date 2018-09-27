package tree

/*size of a tree is the number of elements in the tree*/

func Size(n *Node) int {
	if n == nil {
		return 0
	}
	return Size(n.Right) + Size(n.Left) + 1
}
