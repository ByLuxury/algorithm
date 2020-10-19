package leetcode

import "testing"

// TestCutRope ...
func TestCutRope(t *testing.T) {
	input := 10
	expected := 36
	got := cuttingRope(input)
	if got != expected {
		t.Fatalf("test failed,expected:%v,got:%v", expected, got)
	}
}
