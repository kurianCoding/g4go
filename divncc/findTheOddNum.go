package divncc

/*this function uses binary search to find the unique number
which occurs only onces in the given array*/

func rOddNum(array []int, mid, val int) int {
	return 0
}

func OddNum(array []int) int {

	return rOddNum(array, len(array)/2, array[len(array)/2])
}
