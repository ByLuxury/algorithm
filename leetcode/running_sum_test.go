package leetcode

import "testing"

// RunningSumObj ...
type RunningSumObj struct {
	Nums     []int
	Expected []int
}

func TestRunningSum(t *testing.T) {

	cases := []RunningSumObj{

		{
			Nums:     []int{1, 2, 3, 4},
			Expected: []int{1, 3, 6, 10},
		},
		{
			Nums:     []int{1, 1, 1, 1, 1},
			Expected: []int{1, 2, 3, 4, 5},
		},
		{
			Nums:     []int{3, 1, 2, 10, 1},
			Expected: []int{3, 4, 6, 16, 17},
		},
	}

	for i, c := range cases {
		actual := runningSum(c.Nums)
		if !c.equal(actual) {
			t.Fatalf("test failed,case index:%+v,expected:%+v,got:%+v\n", i, c.Expected, actual)
		}

	}

}

func (r *RunningSumObj) equal(actual []int) bool {

	if len(actual) != len(r.Expected) {
		return false
	}
	for i, v := range r.Expected {
		if actual[i] != v {
			return false
		}
	}
	return true
}
