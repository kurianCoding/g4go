package DP

func MaxSubarraySum(a []int) int {
	var maxSumTillNow int = 0
	var SumTillNow int = 0
	for _, val := range a {
		SumTillNow = SumTillNow + val
		if maxSumTillNow < SumTillNow {
			maxSumTillNow = SumTillNow
		}
		if SumTillNow < 0 {
			SumTillNow = 0
		}
	}
	return maxSumTillNow
}
