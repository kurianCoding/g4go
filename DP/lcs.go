package DP

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func Lcs(a, b string, m, n int) int {
	if m == 0 || n == 0 {
		return 0
	}
	if a[m-1] == b[n-1] {
		return 1 + Lcs(a, b, m-1, n-1)
	}
	return max(Lcs(a, b, m, n-1), Lcs(a, b, m-1, n))
}
