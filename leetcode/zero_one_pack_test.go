package leetcode

import (
	"fmt"
	"testing"
)

// TestZeroPack ...
func TestZeroPack(t *testing.T) {
	volume := 10
	c := []int{2,2,6,5,4}
	v := []int{6,3,5,4,6}
	res := zeroOnePackOpt(volume, c, v)
	fmt.Println(res)
}
