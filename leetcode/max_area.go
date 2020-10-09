package leetcode

// 题目描述：
// 给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。
// 在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。
// 找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

// 输入：[1,8,6,2,5,4,8,3,7]
// 输出：49

func maxArea(height []int) int {
	left, right := 0, len(height)-1 // 0,5
	var res int
	for left < right {
		temp := minNumber(height[left], height[right]) * (right - left)
		res = maxNumber(res, temp)
		if height[left] <= height[right] {
			left++
		} else {
			right--
		}
	}

	return res

}

func minNumber(a, b int) int {

	if a > b {
		return b
	}

	return a
}

func maxNumber(a, b int) int {
	if a > b {
		return a
	}

	return b

}

