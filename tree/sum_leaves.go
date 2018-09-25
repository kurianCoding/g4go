package tree

/*
   given a tree, find the the sum of leaves at each level
   having this value, find the product of all these sums
*/

func isLeaf(n *Node) bool {
	return n.Left == nil && n.Right == nil
}
func SumAndProduct(root *Node) int {
	if root == nil {
		return 0
	}
	var result int = 1
	var q = make([]*Node, 0)
	q = append(q, root)
	for {
		var nodeCount int = len(q)
		if nodeCount == 0 {
			break
		}
		var levelSum int
		var leafFound bool = false
		for nodeCount > 0 {
			n := q[len(q)-1]
			if isLeaf(n) {
				leafFound = true
				levelSum = levelSum + n.Val
			}
			q = q[0 : len(q)-1]
			if n.Left != nil {
				q = append(q, n.Left)
			}
			if n.Right != nil {
				q = append(q, n.Right)
			}
			nodeCount = nodeCount - 1
		}
		if leafFound {
			result = result * levelSum
		}
	}
	return result
}
