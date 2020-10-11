package leetcode

import "testing"

// TestGetLeastNumbers ...
func TestGetLeastNumbers(t *testing.T) {

	arr := []int{3,2,1}
	k := 2
	res := getLeastNumbers(arr, k)
	t.Logf("res:%v", res)
}
