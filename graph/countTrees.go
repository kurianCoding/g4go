package graph

/*given an array that represetns a fortest
the function will return an integer which is
the number of trees in the forest*/
func DepthFirstSearchUtil(forest []int, visited map[int]bool, count int) int {
	for key, val := range forest {
		if !visited[key] {
			visited[key] = true
			return count + DepthFirstSearchUtil(forest, visited, val)
		}
	}
}
func CountTrees(forest []int) int {

	return 2
}
