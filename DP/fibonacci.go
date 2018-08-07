package DP

func Fib(a int) int {
	if a <= 1 {
		return a
	}
	return Fib(a-2) + Fib(a-1)
}
