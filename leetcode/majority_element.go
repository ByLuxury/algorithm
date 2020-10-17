package leetcode

// 数组中占比超过一半的元素称之为主要元素。给定一个整数数组，找到它的主要元素。若没有，返回-1。
//
//示例 1：
//
//输入：[1,2,5,9,5,9,5,5,5]
//输出：5

func majorityElement(nums []int) int {

	if len(nums) == 0 {
		return -1
	}
	//摩尔投票法解决
	res := -1
	vote := 0

	for _,v := range nums {
		if vote == 0 {
			res = v
		}

		if res == v {
			vote ++
		}else {
			vote --
		}
	}

	// 2 次遍历验证是否超过一半
	var count int
	t := len(nums)/2+1
	for _,v := range nums {
		if v == res {
			count ++
		}
		if count == t {
			return v
		}
	}

	return -1
}
