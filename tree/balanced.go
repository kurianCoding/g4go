package tree

/*
    Balanced tree:
	A tree is called balanced if the difference between the left and right
	sub tree is less than one at any level
*/

/* this function takes as input the root node of a tree and checks if it is
balanced or not*/

/*
   to check if a tree is balanced, check the height of the left subtree and
   right subtree, and check if their difference is more than one, if its not
   then the tree is balanced
*/

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -1 * a
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func height(n *Node) int {
	if n == nil {
		return 0
	}
	return 1 + max(height(n.Left), height(n.Right))
}
func Balanced(n *Node) bool {
	if n == nil {
		return true
	}
	lh := height(n.Left)
	rh := height(n.Right)
	if abs(lh-rh) <= 1 && Balanced(n.Left) && Balanced(n.Right) {
		return true
	}

	return false
}
