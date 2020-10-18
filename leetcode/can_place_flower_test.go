package leetcode

import "testing"

// TestCanPlaceFlower ...
func TestCanPlaceFlower(t *testing.T) {

	cases := []struct {
		InputArray []int
		InputN     int
		Expected   bool
	}{
		{
			InputArray: []int{1},
			InputN:     0,
			Expected:   true,
		},
		{
			InputArray: []int{1, 0, 0, 0, 1},
			InputN:     1,
			Expected:   true,
		}, {
			InputArray: []int{1, 0, 0, 0, 1},
			InputN:     2,
			Expected:   false,
		},
	}

	for i, c := range cases {
		actual := canPlaceFlowers(c.InputArray, c.InputN)
		if actual != c.Expected {
			t.Fatalf("test failed,index:%v,ecpected:%v,got:%v", i, c.Expected, actual)
		}
	}
}
