package leetcode

// 三角形最小路径和 动态规划解决
// minimumTotal
func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 || len(triangle[len(triangle)-1]) == 0 {
		return  0
	}

	dp := make([]int, len(triangle[len(triangle)-1]))

	for i := 0; i < len(triangle[len(triangle)-1]); i++ {
		dp[i] = triangle[len(triangle)-1][i] //赋值最后一行数据
	}
	// 自底向上
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			dp[j] = getMinValue(dp[j], dp[j+1]) + triangle[i][j]
		}
	}
	return dp[0]
}

func getMinValue(a, b int) int {

	if a < b {
		return a
	}

	return b
}
