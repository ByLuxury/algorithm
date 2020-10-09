package leetcode

import (
	"fmt"
	"testing"
)

// TestCoinChange ...
func TestCoinChange(t *testing.T) {
	coins := []int{1, 2, 5}
	amount := 11
	res := coinChange(coins, amount)
	fmt.Println(res)
}
