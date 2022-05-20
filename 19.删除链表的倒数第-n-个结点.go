/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第 N 个结点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	tail := head
	log := head
	for i := 1; i < n; i++ {
		tail = tail.Next
	}
	if tail.Next == nil {
		return head.Next
	}else{
		tail = tail.Next
	}
	for tail.Next != nil {
		tail = tail.Next
		head = head.Next
	}
	head.Next = head.Next.Next
	return log
}
// @lc code=end

