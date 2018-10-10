package tree

/*
   post order transversal means printing the left subtree
   then print the right subtree and then finally print
   the root
*/
func search(arr []int, x, n int) int {
	/*
		return the index
	*/
	for i := 0; i < len(arr); i++ {
		if arr[i] == x {
			return i
		}
	}
	return -1
}
func PrintPostOrder(pre, in []int, res *[]int) {
	/*
	   this function takes in pre order and in order transversals
	   of the tree and returns a slice which is the post order
	   transversal.
	*/
	root := search(in, pre, n)
	if root != 0 {
		PrintPostOrder(in, pre+1, root, res)
	}
	if root != n-1 {
		PrintPostOrder(in+root, pre+root, n-root-1, res)
	}
	(*res) = append((*res), pre[0])
	return
}
