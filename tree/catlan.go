package tree

/*catlan numbers can be used to find the number of
trees possible with n given nodes*/
func Catlan(n int) int {
	if n <= 1 {
		return 1
	}
	var res int
	for i := 0; i < n; i++ {
		res = res + Catlan(i)*Catlan(n-1-i)
	}
	return res
}
func PCatlan(n int) int {
	var f int = 1
	for i := 1; i <= 2*n; i++ {
		f = f * i
	}
	var k int = 1
	for j := 1; j <= n+1; j++ {
		k = k * j
	}
	var h int = 1
	for l := 1; l <= n; l++ {
		h = h * l
	}
	return f / (k * h)
}
