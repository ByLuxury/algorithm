package leetcode

import "testing"

// TestFindNumberIn2DArray ...
func TestFindNumberIn2DArray(t *testing.T) {
	input := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	target := 5
	result := findNumberIn2DArray(input,target)

	t.Log(result)
}
