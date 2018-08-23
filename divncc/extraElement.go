package divncc

/*given 2 arrays with sorted elements, the goal is to find that element
which occurs in one of the arrays but not in the other*/

/* this function takes int two arrays and returns the extra element
   it executes O(n)
*/
func ExtraElement(a []int, b []int) int {
	res := 0
	if len(a) < len(b) {
		for key, val := range a {
			if val != b[key] {
				return b[key]
			}
		}
		res = b[len(b)-1]
	} else {
		for key, val := range b {
			if val != a[key] {
				return a[key]
			}
		}
		res = a[len(a)-1]
	}
	return res
}
