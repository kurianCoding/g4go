package DP

func BinomialCoeff(a, b int) int {
	if a < b {
		return 0
	}
	if b == 0 || a == b {
		return 1
	}
	return BinomialCoeff(a-1, b-1) + BinomialCoeff(a-1, b)
}
