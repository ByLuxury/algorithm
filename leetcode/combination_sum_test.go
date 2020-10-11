package leetcode

import (
	"fmt"
	"testing"
)

// TestCombinationSum ...
func TestCombinationSum(t *testing.T)  {

	candidates := []int{2,6,3,7}
	target := 7
	res := combinationSum(candidates,target)
	fmt.Println(res)

}
