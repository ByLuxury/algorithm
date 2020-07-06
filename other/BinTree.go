package other

import "fmt"

type BinTree struct {
	data      int
	leftNode  *BinTree
	rightNode *BinTree
}

func insetData(root *BinTree, data int) {
	if root != nil {
		if data < root.data {
			if root.leftNode == nil {
				root.leftNode = &BinTree{data: data}
			} else {
				insetData(root.leftNode, data)
			}
		} else {
			if root.rightNode == nil {
				root.rightNode = &BinTree{data: data}
			} else {
				insetData(root.rightNode, data)
			}
		}
	}

}

//preTraversal binTree
func preTraverval(root *BinTree) {
	if root != nil {
		fmt.Print(root.data, "-")
		preTraverval(root.leftNode)
		preTraverval(root.rightNode)
	}
}

//midTraversal bintree
func midTraversal(root *BinTree) {
	if root != nil {
		midTraversal(root.leftNode)
		fmt.Print(root.data, "-")
		midTraversal(root.rightNode)
	}
}

//finalTraversal binTree
func finalTraversal(root *BinTree) {
	if root != nil {
		finalTraversal(root.leftNode)
		finalTraversal(root.rightNode)
		fmt.Print(root.data, "-")
	}
}
