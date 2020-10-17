package leetcode

// 给定一个整数数组和一个整数 k，判断数组中是否存在两个不同的索引 i 和 j，使得 nums [i] = nums [j]，并且 i 和 j 的差的 绝对值 至多为 k。
// 示例:
// 输入: nums = [1,2,3,1], k = 3
// 输出: true
// 思路：使用hash
func containsNearbyDuplicate(nums []int, k int) bool {

	m := make(map[int]int)
	for i, v := range nums {
		if _,ok := m[v];ok {
			return true
		}

		m[v]=v
		if len(m)>k {
			delete(m,nums[i-k])
		}

	}
	return false
}
