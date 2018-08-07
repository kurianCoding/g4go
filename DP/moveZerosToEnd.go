/* this function moves zeros to the end of the array and
   returns the array*/
package DP

func CleanZeros(array []int) []int {
	var result []int
	var end []int
	for _, val := range array {
		if val != 0 {
			result = append(result, val)
		} else {
			end = append(end, val)
		}
	}
	return append(result, end...)
}
