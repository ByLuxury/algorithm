package leetcode

import (
	"testing"
)

// TestFlipAndInvertImage ...
func TestFlipAndInvertImage(t *testing.T) {

	input := [][]int{
		{1, 1, 0}, {1, 0, 1}, {0, 0, 0},
	}

	res := flipAndInvertImage(input)
	t.Logf("res:%+v", res)
}
