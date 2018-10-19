package bst

import (
	"geeksforgeeks/tree"
)

/*
   lowest common ancestor is the nearest
   ancestor to two nodes

   this function works on the principle that the
   lca is found while traversing the tree such that
   the value of the node lies between the two nodes.
*/

func LowestCommonAncestor(root, a, b *tree.Node) *tree.Node {
	/*
		this function takes in a and b nodes and returns
		the lowest common ancestor node.
	*/
	if root == nil {
		return nil
	}
	if root.Val > a.Val && root.Val > b.Val {
		return LowestCommonAncestor(root.Right, a, b)
	}
	if root.Val < a.Val && root.Val < b.Val {
		return LowestCommonAncestor(root.Left, a, b)
	}
	return root
}
