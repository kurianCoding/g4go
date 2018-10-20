package tree

/*
   given a binary tree, convert it in to a binary
   search tree, binary search tree, has the pro-
   perty that every subtree has greater value on
   the right and lesser value on the left
*/
func doInOrderTraversal(preArray *[]int, root *Node) {
	if root == nil {
		return
	}
	(*preArray) = append((*preArray), root.Val)
	doInOrderTraversal(preArray, root.Right)
	doInOrderTraversal(preArray, root.Left)
	return
}
func Sort(array *[]int, length int) {
	for i := 0; i < length; i++ {
		for j := i; j < length; j++ {

			if (*array)[i] > (*array)[j] {
				var temp int
				temp = (*array)[j]
				(*array)[j] = (*array)[i]
				(*array)[i] = temp
			}
		}
	}
	return
}
func MakeTree(root *Node, array []int, count int) {
	if root == nil {
		return
	}
	MakeTree(root.Left, array, count+1)
	root.Val = array[count]
	MakeTree(root.Right, array, count+1)
	return
}
func BinaryTree2SubTree(tree *Node) *Node {
	/*
		first we traverse the tree in order traversal,
		then we do a sorting on the array that results from the traversal
		then we build a new tree from the array.
		return the root node of the tree
	*/
	inOrderArray := make([]int, 0)
	doInOrderTraversal(&inOrderArray, tree)
	Sort(&inOrderArray, len(inOrderArray))
	root := NewNode(inOrderArray[0])
	MakeTree(root, inOrderArray, 1)
	return root
}
