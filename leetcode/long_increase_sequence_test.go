package leetcode

import (
	"fmt"
	"testing"
)

func TestLengthOfLIS(t *testing.T) {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}

	res := LengthOfLIS(nums)

	fmt.Println(res)
}
