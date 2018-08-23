package divncc

func CountRot(array []int) int {
	//given is an array of sorted numbers that have been rotated
	//right. This function returns how many times the array has been
	//rotated.
	if array[1] < array[0] {
		return 0
	}
	lenar := len(array)
	for i := 1; i < lenar; i++ {
		if array[i] < array[i-1] {
			return i
		}
	}
	return lenar
}
