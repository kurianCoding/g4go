package bst

/*
   find the single successor using, min and max
*/
func MinMaxSingleTree(pre []int) bool {
	a := pre[len(pre)-1]
	b := pre[len(pre)-2]

	if b > a {
		temp := a
		a = b
		b = temp
	}
	for i := 0; i < len(pre); i++ {
		if pre[i] < a && pre[i] > b {
			/*
				value lies between the minimum and maximum of
				the binary tree
			*/
			return false
		}
	}

	return true
}
