package DP

func HCF(a, b int) int {
	if b > a {
		temp := a
		a = b
		b = temp
	}
	if a%b == 0 {
		return b
	}
	return HCF(b, a%b)
}
