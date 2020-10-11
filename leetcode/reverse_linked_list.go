package leetcode

// reverseList 反转链表
// 输入: 1->2->3->4->5->NULL
// 输出: 5->4->3->2->1->NULL
func reverseList(head *ListNode) *ListNode {

	var  pre *ListNode
	cur := head
	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}
	return pre
}
