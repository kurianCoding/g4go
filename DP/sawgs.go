package DP

/*
   given an array aim is to find the first subarray which
   has the given sum
*/

func getSumArray(array []int, sum int) (int, int) {
	for i := 0; i < len(array); i++ {
		for j := i; j < len(array); j++ {
			if sums(array, i, j) == sum {
				return i, j
			}
		}
	}
	return -1, -1
}

func sums(arr []int, a, b int) int {
	var suma int
	if a == b {
		return a
	}
	for i := a; i < b; i++ {
		suma = suma + arr[i]
	}
	return suma
}
