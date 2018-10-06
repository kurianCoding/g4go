package bst

import (
	"geeksforgeeks/tree"
	"testing"
)

/*
   binary search tree is a way of classifying information,
   all the right nodes of of a subtree contains values greater
   than the node value, naturally the left nodes contain all the
   nodes containing lesser values
*/
func TestLowestCommonAncestor(t *testing.T) {
	/*
		lowest common ancestor takes in three nodes
		it returns one node which is the common
		ancestor
	*/
}

func TestIsLevelTrans(t *testing.T) {
	/*
		the level transversal funtion accpets an array and returns
		a boolean
	*/
}
func TestClosestNeighbourSearch(t *testing.T) {
	c := ClosestNeigbourSearch(root, 5)
	root := tree.NewNode(2)
	a := tree.NewNode(2)
	a1 := tree.NewNode(3)
	a2 := tree.NewNode(1)
	a3 := tree.NewNode(5)
	a4 := tree.NewNode(4)
	root.AddRight(a)
	a.AddRight(a1)
	a.AddLeft(a3)
	a3.AddRight(a4)
	if c.Val != 5 {
		t.Errof("expected %d got %d", 5, c.Val)
	}
}
