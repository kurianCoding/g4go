package DP

func LenProf(a []int, len int) int {
	if len <= 0 {
		return 0
	}
	max_val := 0
	for i := 0; i < len; i++ {
		max_val = max(max_val, a[i]+LenProf(a, len-i-1))
	}
	return max_val
}
