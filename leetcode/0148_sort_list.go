package main

import "fmt"

type ListNode0148 struct {
	Val  int
	Next *ListNode0148
}

func merge0148(a, b *ListNode0148) *ListNode0148 {
	dummy := new(ListNode0148)
	current := dummy
	for a != nil && b != nil {
		if a.Val <= b.Val {
			current.Next = a
			a = a.Next
		} else {
			current.Next = b
			b = b.Next
		}
		current = current.Next
	}
	if a == nil {
		current.Next = b
	} else {
		current.Next = a
	}
	return dummy.Next
}

// 标准解法：归并排序
func sortList(head *ListNode0148) *ListNode0148 {
	// 0.递归结束条件
	if head == nil || head.Next == nil {
		return head
	}
	// 1.中点切断
	slow := head
	previous := (*ListNode0148)(nil)
	for fast := head; fast != nil && fast.Next != nil; fast = fast.Next.Next {
		previous = slow
		slow = slow.Next
	}
	previous.Next = nil
	// 2.递归切断
	left := sortList(head)
	right := sortList(slow)
	// 3.递归合并
	return merge0148(left, right)
}

// 思路总结：在排序中，如果可随机存取就用快速排序，不能的话就用归并排序

func main() {
	a1, a2, a3, a4 := new(ListNode0148), new(ListNode0148), new(ListNode0148), new(ListNode0148)
	a1.Val, a2.Val, a3.Val, a4.Val = 4, 2, 1, 3
	a1.Next, a2.Next, a3.Next = a2, a3, a4
	for p := sortList(a1); p != nil; p = p.Next {
		fmt.Print(p.Val, " ")
	}
}
