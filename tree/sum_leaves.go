package tree

/*
   given a tree, find the the sum of leaves at each level
   having this value, find the product of all these sums
*/

func SumLeaves(t *Tree) int {
	r := t.GetRoot()
	sum := rSum(r)
}
