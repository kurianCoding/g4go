package bst

/*
   single successor, thei program takes the pre order transversal of a tree,
   and by using algorithm, returns if every node in the tree has a single
   decendent of multiple decendants.
*/
func IsSingle(list []int, size int) bool {
	for i := 0; i < len(list)-1; i++ {
		firstDiff := list[i] - list[i+1]
		/*
			the succesive element is either greater
			or lesser
		*/
		secondDiff := list[i] - list[size-1]
		/*
			list[size-1] is the last element
		*/
		if firstDiff*secondDiff < 0 {
			return false
		}
	}
	return true
}
