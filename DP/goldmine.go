package DP

// this function returns the maximum amoutn of gold
//which can be successfully retrived from the mine while
//traversing it
// input is hte array, output is the int which represents
// sum of all the gold that he collected from the mines
func netGold(sum [][]int, length int, breadth int) int {
	/*
		i, j := 0, 0
		for _, row := range array {
			for _, num := range row {
				if i == length || j == breadth {
					continue
				} else {
					sum[i][j] = sum[i][j] + max(array[i][j+1], array[i+1][j])
				}
				j = j + 1

			}
			i = i + 1
		}
		return sum[length-1][breadth-1]*/
	return -1
}
