package bst

/*
   Given two arrays, this program is to decide if the two arrays
   will produce the same binary search tree. This involves, seeing
   if all the subtrees of the two binary search tree will be same

   This program does so by recursively finding the left node and the
   right node and checkin if they are equal.
*/

func rIsSameTree(a, b []int, indl, indh, max, min int) bool {
	/*
	   variables to store the indices of the
	   two search results
	*/
	var i int
	var j int
	for i = indl; i < len(a); i++ {
		if a[i] > min && a[i] < max {
			break
		}
	}
	for j = indh; j < len(b); j++ {
		if b[j] > min && b[j] < max {
			break
		}

	}
	/*
	   search for the value that lies between the maximum and the
	   minimum
	*/
	if i == len(a) && j == len(b) {
		return true
	}
	/*
	   if the search has reached the leaves
	*/
	if i == len(b)^j == len(a) || a[i] != b[j] {
		return false
	}
	/*
		    if either one of the two trees terminate without the other
		    ending.
			or
		    either of the elements are not the same
	*/

	return rIsSameTree(a, b, i+1, j+1, a[i], max) && rIsSameTree(a, b, i+1, j+1, min, a[i])

}
func IsSameTree(a, b []int) bool {
	return rIsSameTree(a, b, 0, 0, 0, 127)
}
