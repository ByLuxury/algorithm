package leetcode

import (
	"fmt"
	"testing"
)

// TestTriangleMinSum ...
func TestTriangleMinSum(t *testing.T) {
	input := [][]int{{-1}, {2, 3}, {1, -1, -3}}
	res := minimumTotal(input)
	fmt.Println(res)

}
