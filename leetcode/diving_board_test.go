package leetcode

import "testing"

// TestDivingBoard ...
func TestDivingBoard(t *testing.T) {

	shorter := 1
	longer := 1
	k := 0
	res := DivingBoard(shorter, longer, k)
	t.Log(res)
}
