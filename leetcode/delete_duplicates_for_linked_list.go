package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。
//
//示例 1:
//
//输入: 1->1->2
//输出: 1->2
//示例 2:
//
//输入: 1->1->2->3->3
//输出: 1->2->3
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list

func deleteDuplicates(head *ListNode) *ListNode {
	cur := head

	for cur != nil && cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}
