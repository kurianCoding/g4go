package bst

import "geeksforgeeks/tree"

/*
   find the largest or equal to element which is lesser than
   the requested key.
*/

func rClosestNeighbourSearch(root, res *tree.Node, diff, val int) *tree.Node {
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

func ClosestNeighbourSearch(root *tree.Node, val int) *tree.Node {
	return rClosestNeighbour(root, root, root.Val-val, nil)
}
