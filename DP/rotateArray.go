package DP

/* rotate array by n, given an array, rotate it by n
   return the rotated array.*/
func RotateArray(array []int, k int) []int {
	var result []int
	length := len(array)
	for i := 0; i < length; i++ {
		result = append(result, array[(k+i+1)%length])
	}
	return result
}

/* adding one because the indexing in array puts it one index back
   this was not clear from the logic. But it is true.*/
