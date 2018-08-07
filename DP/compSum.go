package DP

func CompleteSumR(array []int, n, sum int) bool {
	if sum == 0 {
		return true
	}
	if sum != 0 && n == 0 {
		return false
	}
	if array[n-1] > sum {
		return CompleteSumR(array, n-1, sum)
	}

	return CompleteSumR(array, n-1, sum) || CompleteSumR(array, n-1, sum-array[n-1])
}
func CompleteSum(array []int, sum int) bool {
	return CompleteSumR(array, len(array), sum)
}
