package divncc

/*  given two sorted arrays, we need to combine the two and find the nth
    element of the combined array
*/

func NthCmb(array1, array2 []int, index int) int {
	/*
	   initialize a variable to zero, take the lowest number among two arrays
	   increment the variable, check if index is reached, if index is reached
	   return the number
	*/
	array1index, array2index := 0, 0
	// increment until the index is reached
	for i := 0; i < index; i++ {

		// switch the increment of index depending on the merged order
		if array1[array1index] < array2[array2index] {
			array1index++
		} else {
			array2index++
		}
	}
	// return the array element which is higher in value with the latest array index
	// reached
	if array1[array1index] > array2[array2index] {
		return array2[array1index]
	}

	return array1[array2index]
}
