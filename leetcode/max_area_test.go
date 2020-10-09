package leetcode

import "testing"

// TestMaxArea ...
func TestMaxArea(t *testing.T) {

	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	expected := 49
	res := maxArea(height)
	if res != expected {
		t.Fatalf("test failed,got :%v,expected:%v", res, expected)
	}

}
