package DP

/*
   we have a set of dominoes, with which we are suppoused to tile
   a given dimension. this function finds out how many ways are
   there to do that. It is given that the width of area is 3 so
   dimensions are in the form of 3x(length). INput to the function
   is length, return value is an integer which is the number of
   ways to arragne the dominoes.
*/
func CountWays(length int) int {
	A := make([]int, length+1)
	B := make([]int, length+1)
	A[0] = 1
	A[1] = 0
	B[0] = 0
	B[1] = 1
	for i := 2; i <= length; i++ {
		A[i] = A[i-2] + 2*B[i-1]
		B[i] = A[i-1] + B[i-2]
	}
	return A[length]
}
