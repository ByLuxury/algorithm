package leetcode

// cuttingRope 一根绳子剪成k段，求m[0]*m[1]*...[k-1]段最大乘积，结果取模 1000000007
func cuttingRope(n int) int {
	if n<4 {
		return n-1
	}
	if n == 4 {
		return n
	}

	res := 1
	for n>4 {
		res *= 3
		res %= 1000000007
		n-=3
	}

	return res *n% 1000000007

}