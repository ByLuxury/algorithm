package leetcode

func runningSum(nums []int) []int {

	result := make([]int,0)
	sum := 0
	for _,v := range nums {
		sum +=v
		result = append(result,sum)
	}
	return result
}
