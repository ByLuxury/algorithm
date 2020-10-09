package leetcode

import "testing"

//输入：1->2->4, 1->3->4
//输出：1->1->2->3->4->4

// TestMergeTwoLists ...
func TestMergeTwoLists(t *testing.T) {

	// 构造两个链表并插入测试数据
	l1 := new(ListNode)
	l2 := new(ListNode)
	l1.Val = 1
	l1.Next = &ListNode{
		Val: 2,
		Next: &ListNode{
			Val:  4,
			Next: nil,
		},
	}

	l2.Val = 1
	l2.Next = &ListNode{
		Val: 3,
		Next: &ListNode{
			Val:  4,
			Next: nil,
		},
	}

	res := mergeTwoLists(l1, l2)

	for res != nil {
		t.Logf("%v ", res.Val)
		res = res.Next
	}
}
