package leetcode

func findSingleNumber(input []int) int {

	if len(input) == 0 {
		return -1

	}
	res := 0

	for _, v := range input {
		res = res ^ v
	}

	return res
}


