package leetcode

// 题目描述：
// 有 N 件物品和一个容量为 V 的背包。每种物品都有自己的容量（c[i]）和价值（v[i]）
// 在限定的总容量（V）内，选择其中若干个（也即每种物品可以选0个或1个）
// 设计选择方案使得物品的总价值最高，求出最大总价值。
// 思路：动态规划
// 状态转移方程：

func zeroOnePackOpt(volume int, c []int, value []int) int {

	if volume == 0 || len(c) == 0 || len(value) == 0 {
		return 0
	}

	dp := make([]int, volume+1)

	n := len(c)
	for i := 0; i < n; i++ {
		for j := volume; j >= c[i]; j-- {
			dp[j] = zeroOnePackOptMax(dp[j], dp[j-c[i]]+value[i])
		}
	}
	return dp[volume]

}

func zeroOnePackOptMax(a, b int) int {

	if a > b {
		return a
	}
	return b
}
