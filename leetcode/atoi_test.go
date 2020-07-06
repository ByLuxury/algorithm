package leetcode

import (
	"testing"
)

// TestAtoi ...
func TestAtoi(t *testing.T) {

	testCase := []MyAtoiCase{
		{
			Input:    "  -44",
			Expected: -44,
		},
		{
			Input:    "H124",
			Expected: 0,
		},
		{
			Input:    "102 this code",
			Expected: 102,
		},
		{
			Input:    "+-12",
			Expected: 0,
		},
		{
			Input:    "1234",
			Expected: 1234,
		},

		{
			Input:    "1234def3",
			Expected: 1234,
		},

		{
			Input:    "-1234def3",
			Expected: -1234,
		},

		{
			Input:    "2147483648",
			Expected: 2147483647,
		},
		{
			Input:    "-91283472332",
			Expected: -2147483648,
		},
		{
			Input:    "18446744073709551617",
			Expected: 2147483647,
		},
	}

	for _, c := range testCase {
		actual := c.MyAtoi(c.Input)
		if actual != c.Expected {
			t.Fatalf("test failed:,expected:%d,got:%d\n", c.Expected, actual)
		}
	}

}
