package leetcode

import "sort"

// 题目描述：
// 给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
// candidates 中的数字可以无限制重复被选取。
// 说明： 所有数字（包括 target）都是正整数。 解集不能包含重复的组合。 

func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	sort.Ints(candidates)

	var dfs func(int, int, []int)
	dfs = func(i int, target int, path []int) {
		if i >= len(candidates) {
			return
		}
		if target < 0 {
			return
		}
		if target == 0 {
			res = append(res, append([]int{}, path...))
			return
		}
		dfs(i+1, target, path)
		dfs(i, target-candidates[i], append(path, candidates[i]))
		return
	}

	dfs(0, target, []int{})
	return res
}

func removeEleFromSlice(s []int, index int) []int {
	s[len(s)-1], s[index] = s[index], s[len(s)-1]
	return s[:len(s)-1]
}

//func remove(slice []int, index int) []int {
//
//	return append(slice[:index], slice[index+1:]...)
//}

//示例 1：
//
//输入：candidates = [2,3,6,7], target = 7,
//所求解集为：
//[
//  [7],
//  [2,2,3]
//]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/combination-sum
