package bst

import "geeksforgeeks/tree"

/*
   binary search trees are classifications of elements as lesser
   or greater, we have to find the nth largest element in the
   tree, this function does  it.

   It traverses first the right and then the left subtree of the
   given tree. If it hits a laef or if the number of visits is
   greater than n then it stops and returns. If it hits the
   correct number of largest, it will return the answer
*/

func NthLargest(root *tree.Node, count, limit int) (bool, *tree.Node) {
	if root == nil || count > limit {
		return false, nil
	}

	if ok, ans := NthLargest(root.right, count, limit); ok == true {
		return true, ans
	}
	count = count + 1
	if count == limit {
		return true, root
	}
	if ok, ans := NthLargest(root.left, count, limit); ok == true {
		return true, ans
	}
}
