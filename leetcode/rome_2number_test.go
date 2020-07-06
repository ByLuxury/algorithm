package leetcode

import (
	"fmt"
	"testing"
)

// TestRome2Number_RomanToInt ...
func TestRome2Number_RomanToInt(t *testing.T) {

	testCase := []Rome2Number{
		{
			Input:    "IV",
			Expected: 4,
		},
		{
			Input:    "IX",
			Expected: 9,
		},
		{
			Input:    "LVIII",
			Expected: 58,
		},
		{
			Input:    "MCMXCIV",
			Expected: 1994,
		},
	}
	for _, c := range testCase {
		actual := c.RomanToInt(c.Input)
		if actual != c.Expected {
			fmt.Printf("test failed:,expected:%d,got:%d\n", c.Expected, actual)
		}
	}

}
