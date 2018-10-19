package bst

import "geeksforgeeks/tree"

/*
   this algorithm searches the binary search algorithm
   by searching the binary tree for a given key, it returns
   the pointer to the object that stores the given key
*/
func SearchBinaryTree(root *tree.Node, val int) *tree.Node {
	if root == nil || root.Val == val {
		return root
	}
	if root.Val < val { /*
			value of the node is lesser than the value we
			are looking for
		*/
		return SearchBinaryTree(root.Right, val)
	}
	return SearchBinaryTree(root.Left, val) /*
		the node value is greater than
		the value we are looking for
	*/
}
