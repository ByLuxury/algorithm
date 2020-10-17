package leetcode

import "testing"

// TestMajorityElement ...
func TestMajorityElement(t *testing.T) {
	cases := []struct {
		Input    []int
		Expected int
	}{
		{
			Input:    []int{1, 2, 5, 9, 5, 9, 5, 5, 5},
			Expected: 5,
		},
		{
			Input:    []int{3, 2},
			Expected: -1,
		},
		{
			Input:    []int{2, 2, 1, 1, 1, 2, 2},
			Expected: 2,
		},
	}

	for i, c := range cases {
		got := majorityElement(c.Input)
		if got != c.Expected {
			t.Fatalf("test failed,expected:%v,got:%v,index:%v", c.Expected, got, i)
		}
	}

}
