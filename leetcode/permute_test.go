package leetcode

import (
	"fmt"
	"testing"
)

// TestPermutation ...
func TestPermutation(t *testing.T) {

	res := []int{1, 2, 3}
	out := Permutation(res)
	for _, v := range out {
		fmt.Println(v)
	}
}
