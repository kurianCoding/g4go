package bst

/*
   given a root of tree node, and an integer K
   find the Kth smallest element of the tree
*/
func KthSmall(root *Node, k int) int {
	/*
		the function takes in the root node, it
		returns an integers which is the nth
		smallest number
	*/
	if root.Right != nil {
		queue = append(queue, root.Right.Val)
	}

}
