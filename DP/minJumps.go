package DP

const (
	MaxInt int = int(^uint(0) >> 1)
)

func MinJumpsR(a []int, h, l int) int {
	if h == l {
		return 0
	}
	if a[h] == 0 {
		return MaxInt
	}
	var min int = MaxInt
	for i := h + 1; i <= l && i <= h+a[h]; i++ {
		jumps := MinJumpsR(a, i, l)
		if jumps != MaxInt && jumps+1 < min {
			min = jumps + 1
		}
	}
	return min
}
func MinJumps(a []int) int {
	return MinJumpsR(a, 0, len(a)-1)
}
