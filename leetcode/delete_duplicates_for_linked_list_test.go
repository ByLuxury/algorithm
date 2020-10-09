package leetcode

import (
	"fmt"
	"testing"
)

// TestDeleteDuplicatesForLinkedList ...
func TestDeleteDuplicatesForLinkedList(t *testing.T) {

	l1 := new(ListNode)
	l1.Val = 1
	l1.Next = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
			},
		},
	}
	res := deleteDuplicates(l1)
	for res != nil {
		fmt.Printf("%v ", res.Val)
		res = res.Next
	}
}
