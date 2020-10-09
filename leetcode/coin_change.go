package leetcode

// 零钱兑换
// 思路：使用动态规划，类比为爬楼梯
// dp[i]表示，爬上i级台阶的最少步数
// 状态转移方程：dp[i] = min{dp[i-conis[j]]}+1
func coinChange(coins []int, amount int) int {

	dp := make([]int, amount+1)

	for i := 1; i <= amount; i++ {
		dp[i] = -1
		for j := 0; j < len(coins); j++ {
			if i < coins[j] || dp[i-coins[j]] == -1 {
				continue
			}

			count := dp[i-coins[j]] + 1
			if dp[i] == -1 || dp[i] > count {
				dp[i] = count
			}

		}

	}

	return dp[amount]
}


// 题目描述：
// 给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。
// 每种硬币的数量是无限的。

// 示例 1：
// 输入：coins = [1, 2, 5], amount = 11
// 输出：3
// 解释：11 = 5 + 5 + 1

// 示例 2：
// 输入：coins = [2], amount = 3
// 输出：-1

// 示例 3：
// 输入：coins = [1], amount = 0
// 输出：0
