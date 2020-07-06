package leetcode

import "testing"

// TestLongestCommonPrefix ...
func TestLongestCommonPrefix(t *testing.T) {
	testCases := []LCP{
		{
			Input:    []string{"dog", "racecar", "car"},
			Expected: "",
		},
		{
			Input:    []string{"flower", "flow", "flight"},
			Expected: "fl",
		},
		{
			Input:    []string{"leetcode", "leet", "lee"},
			Expected: "lee",
		},
		{
			Input:    []string{"", "wish"},
			Expected: "",
		},
	}

	for _, c := range testCases {
		got := c.LongestCommonPrefix(c.Input)
		if got != c.Expected {
			t.Fatalf("test failed,expected:%s,got:%s", c.Expected, got)
		}
	}

}
