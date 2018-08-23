package divncc

/*
   given a sorted array the idea is to find the most repeated element of the array
*/

func IsMajority(array []int, a int) bool {
	//this function checks if the given element is the most repeated element in the array
	countmap := make(map[int]int)
	for _, val := range array {
		countmap[val]++
	}
	tempmax := 0
	tempmaxval := 0
	for key, count := range countmap {
		if count > tempmaxval {
			tempmaxval = count
			tempmax = key
		}
	}
	// if the maximum counted element is equal to the
	// given value
	return tempmax == a
}
