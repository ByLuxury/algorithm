package leetcode

import (
	"fmt"
	"testing"
)

// TestMergeTrees ...
func TestMergeTrees(t *testing.T) {

	t1 := &TreeNode{
		Val: 1,
	}
	t1.Left = &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 5,
		},
	}
	t1.Right = &TreeNode{
		Val: 2,
	}
	t2 := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	res := MergeTrees(t1, t2)
	travelTree(res)

}

func travelTree(tree *TreeNode) {

	if tree != nil {
		fmt.Println(tree.Val)
		travelTree(tree.Left)
		travelTree(tree.Right)
	}

}
