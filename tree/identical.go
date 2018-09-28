package tree

/* function takes in two trees, traverses them, and compares each node
with a node from the other tree to check if both trees are identical
*/
func Identical(n, n1 *Node) bool {
	if n == nil && n1 == nil {
		return true
	}
	if n == nil && n1 != nil || n != nil && n1 == nil {
		return false
	}
	return n.Val == n1.Val && Identical(n.Right, n1.Right) && Identical(n.Left, n1.Left)
}
