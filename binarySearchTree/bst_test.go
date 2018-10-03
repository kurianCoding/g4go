package bst

/*
   binary search tree is a way of classifying information,
   all the right nodes of of a subtree contains values greater
   than the node value, naturally the left nodes contain all the
   nodes containing lesser values
*/

func TestClosestNeighbourSearch(t *testing.T) {
	c := ClosestNeigbourSearch(root, 5)
	if c.Val != 5 {
		t.Errof("expected %d got %d", 5, c.Val)
	}
}
