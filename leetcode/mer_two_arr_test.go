package leetcode

import (
	"fmt"
	"testing"
)

// TestMergeTwoArr ...
func TestMergeTwoArr(t *testing.T) {

	arr1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	arr2 := []int{2, 5, 6}
	n := 3
	mergeTwoArr(arr1,m,arr2,n)
	fmt.Println(arr1)
}
