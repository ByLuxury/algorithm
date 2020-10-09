package leetcode

//	最长上升子序列
//	输入: [10,9,2,5,3,7,101,18]
//	输出: 4
//	解释: 最长的上升子序列是 [2,3,7,101]，它的长度是 4。

// 	思路：使用动态规划,
// 	状态：dp[i] 表示走到i位置时的长度
//  状态转移方程：dp[i]=MAX{dp[j]}+1 (0<=j<i,a[j]<a[i])
//
// LengthOfLIS ...
func LengthOfLIS(nums []int) int {

	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	// 默认长度为1
	for i := 0; i < len(nums); i++ {
		dp[i] = 1

	}
	res := 0
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = getMax(dp[i], dp[j]+1)
			}
		}
		res = getMax(res, dp[i])

	}
	return res
}

func getMax(a, b int) int {

	if a > b {
		return a
	}
	return b
}
