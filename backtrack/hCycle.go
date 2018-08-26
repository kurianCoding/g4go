package backtrack

/*
   A hamiltonian cycle is a path from first node, that touches
   all other nodes only once and returns to the first node

   this function takes in a incidence array(2-D) and returns a boolean.
   If there exists a hamiltonian path then it returns true and if it doesn't
   it returns false.
*/
/* solution:
   create an empty array, add the first node, before adding the next
   node check if that node is added to path or if its neighbors have
   been added to the path. If yes then skip the node, if no then add it
*/

func isSafeHM(i int, g [][]bool, path []int, length, pos int) bool {

	if g[path[pos-1]][i] == false {
		return false
	}

	for j := 0; j < length; j++ {
		if path[j] == 0 {
			return false
		}
	}
	return true
}

func checkHM(g [][]bool, path []int, pos, length int) bool {

	if pos == length {
		if g[path[pos-1]][path[0]] == true {
			return true
		} else {
			return false
		}
	}

	for i := 0; i < length; i++ {
		if isSafeHM(i, g, path, length, pos) {
			path[pos] = i
			if checkHM(g, path, pos+1, length) == true {
				return true
			}
		}
	}
	return false
}

func hasHM(g [][]bool, length int) bool {
	path := make([]int, length)
	for i := 0; i < length; i++ {
		path[i] = -1
	}
	path[0] = 0
	if checkHM(g, path, 1, length) {
		return true
	}
	return false
}
