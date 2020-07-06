package other

import (
	"testing"
)

// TestSearchData 。。。
func TestSearchData(t *testing.T) {
	var a [LEN]int
	//var a [LEN]int=[LEN]int{1,2,3,4,5,6,7,8,9,7}
	for i := 0; i < len(a); i++ {
		a[i] = i + 1
	}
	a[len(a) - 1] = a[6]
	result := search(a)
	t.Log(result)
}
