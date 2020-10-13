package leetcode

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T)  {
	arr := []int{1, 4, 6, 7, 9, 12, 43, 45, 65, 122, 211}
	low, high := 0, len(arr)-1
	res := BinarySearch(arr, low, high, 6)
	fmt.Println(res)

}
