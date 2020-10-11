package leetcode

import "testing"

func TestQuickSort(t *testing.T) {

	nums := []int{5, 1, 1, 2, 0, 13, 4, 32, 9}
	sortArray(nums)
	t.Logf("result:%+v", nums)
}
