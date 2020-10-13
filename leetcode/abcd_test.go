package leetcode

import (
	"testing"
)

func TestABCD(t *testing.T) {

	arr := []int{1, 2, 3, 4, 5, 6}
	res := calcSum(arr)

	t.Logf("res:%+v", res)
}
