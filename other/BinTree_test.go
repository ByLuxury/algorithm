package other

import (
	"fmt"
	"testing"
)

// TestBinTree ...
func TestBinTree(t *testing.T) {
	x := []int{12, 76, 35, 22, 16, 48, 90, 46, 9, 40};
	//var root *BinTree
	root := &BinTree{data: x[0]}
	//root.data = x[0]
	root.leftNode = nil
	root.rightNode = nil
	for i := 1; i < len(x); i++ {
		insetData(root, x[i])
	}
	fmt.Println("preTraversal:")
	preTraverval(root)
	fmt.Println()
	fmt.Println("midTraversal:")
	midTraversal(root)
	fmt.Println()
	fmt.Println("finalTraversal:")
	finalTraversal(root)
}
