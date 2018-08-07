package DP

/*
	the idea of cutting a rod, to pices, keeping the
	price of each pice in mind. The idea is to maximize
	the price.
*/
func OptRate(priceArray []int, rodLength int) int {
	return backtrackOptRate(priceArray, rodLength)
}
func maxOptRate(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func backtrackOptRate(priceArray []int, rodLength int) int {
	/*
		this function takes rod length and price array as
		inputs, it gives back the maximum price which
		can be obtained as the result by optimizign the
		price at each step and recursively calculating the
		maximum price
	*/
	var maxVal int = 0
	if rodLength < 0 {
		return maxVal
	}
	for i := 0; i < rodLength; i++ {
		maxVal = max(maxVal, backtrackOptRate(priceArray, rodLength-i-1)+priceArray[i])
	}
	return maxVal
}
