package leetcode

import "testing"

// TestContainNearbyDuplicate ...
func TestContainNearbyDuplicate(t *testing.T) {

	arr := []int{1, 2, 3, 1}
	k := 3
	res := containsNearbyDuplicate(arr, k)
	if !res  {
		t.Fatal("test failed,expected[true],got[false]")
	}
}
