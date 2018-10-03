package bst

/*
   find the largest or equal to element which is lesser than
   the requested key.
*/

func rClosestNeighbourSearch(root, res *Node, diff, val int) *Node {
	if root == nil {
		return res
	}
	if diff > root.Val-val {
		diff = root.Val - val
		res = root
	}
	if val > root.Val {
		return rClosestNeighbourSearch(root.Right, res, diff, val)
	}
	return rClosestNeighbourSearch(root.Left, res, diff, val)
}

func ClosestNeighbourSearch(root *Node, val int) *Node {
	return rClosestNeighbour(root, root, root.Val-val, nil)
}
