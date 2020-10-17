package leetcode
// 给你两个有序整数数组 nums1 和 nums2，请你将 nums2 合并到 nums1 中，使 nums1 成为一个有序数组。
//说明：
//初始化 nums1 和 nums2 的元素数量分别为 m 和 n 。
//你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。

// 输入：
//nums1 = [1,2,3,0,0,0], m = 3
//nums2 = [2,5,6],       n = 3
//
//输出：[1,2,2,3,5,6]

// mergeTwoArr ...
func mergeTwoArr(nums1 []int, m int, nums2 []int, n int) {

	// 双指针
	copyNums1 := make([]int, m)

	copy(copyNums1, nums1[:m])
	p1 := 0
	p2 := 0

	p := 0
	for p1 < m && p2 < n {
		if copyNums1[p1] < nums2[p2] {
			nums1[p] = copyNums1[p1]
			p1++
		} else {
			nums1[p] = nums2[p2]
			p2++
		}
		p++
	}
	if p1 < m {
		for i := p1+p2; i < m+n; i++ {
			nums1[i] = copyNums1[p1]
			p1++
		}
	}

	if p2 < n {

		for i := p1+p2; i < m+n; i++ {
			nums1[i] = nums2[p2]
			p2++
		}
	}
}
