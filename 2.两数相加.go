/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
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
type ListNode struct {
	Val int
	Next *ListNode
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	list1 := l1
	list2 := l2
	ret := new(ListNode)
	inc := 0
	var v1, v2 int = list1.Val, list2.Val
	p := ret; 
	for  {
		sum := v1 + v2 + inc
		p.Val = sum % 10
		inc = int(sum / 10)
		
		if list1.Next != nil || list2.Next != nil || inc != 0{
			p.Next = new(ListNode)
			p = p.Next
		}else{
			break
		}

		if list1.Next != nil{
			list1 = list1.Next
			v1 = list1.Val
		}else{
			v1 = 0
		}

		if list2.Next != nil{
			list2 = list2.Next
			v2 = list2.Val
		}else{
			v2 = 0
		}

		
	}

	return ret
}
// @lc code=end

