package main

import "fmt"

type ListNode0025 struct {
	Val  int
	Next *ListNode0025
}

// 标准解法：双指针固定距离 + 反转链表
func reverseList0025(head *ListNode0025) *ListNode0025 {
	if head == nil || head.Next == nil {
		return head
	}
	previous, current, next := (*ListNode0025)(nil), head, head.Next
	for current != nil {
		// 1.反转
		next = current.Next
		current.Next = previous
		// 2.更新指针
		previous = current
		current = next
	}
	return previous
}

func reverseKGroup(head *ListNode0025, k int) *ListNode0025 {
	dummy := &ListNode0025{Next: head}
	fast := dummy
	for fast != nil {
		// 1.初始化双指针
		previous := fast
		for i := 0; i < k; i++ {
			fast = fast.Next
			if fast == nil {
				return dummy.Next
			}
		}
		// 2.反转
		next := fast.Next                              // 红头
		fast.Next = nil                                // 切断
		fast = previous.Next                           // 更新 fast
		previous.Next = reverseList0025(previous.Next) // 左连接
		fast.Next = next                               // 右连接
	}
	return dummy.Next
}

// 思路总结：
// 1.多指针的情况下，一定要有基准指针，其它指针伴随更新，不然容易出现逻辑漏洞
// 2.反转链表一定要先切断主链表

func main() {
	a1, a2, a3, a4, a5 := new(ListNode0025), new(ListNode0025), new(ListNode0025), new(ListNode0025), new(ListNode0025)
	a1.Val, a2.Val, a3.Val, a4.Val, a5.Val = 1, 2, 3, 4, 5
	a1.Next = a2
	a2.Next = a3
	a3.Next = a4
	a4.Next = a5
	for p := reverseKGroup(a1, 2); p != nil; p = p.Next {
		fmt.Print(p.Val, " ")
	}
}
