package main

import "fmt"

type ListNode0234 struct {
	Val  int
	Next *ListNode0234
}

// 标准解法：快慢指针 + 链表反转法，常规做法是用栈
func isPalindrome(head *ListNode0234) bool {
	if head == nil {
		return false
	}
	// 1.指针移向中间（中间或右一）
	slow := head
	for fast := head; fast != nil && fast.Next != nil; fast = fast.Next.Next {
		slow = slow.Next
	}
	// 2.反转右链表
	previous, current, next := (*ListNode0234)(nil), slow, slow.Next
	for current != nil {
		next = current.Next
		current.Next = previous
		previous = current
		current = next
	}
	// 3.遍历对比
	a := head
	for b := previous; b != nil; a, b = a.Next, b.Next {
		if a.Val != b.Val {
			return false
		}
	}
	return true
}

// 思路总结：
// 1.如果把链表从两侧向中间靠拢遍历，就可以进行比对
// 2.要让右半段能从尾走到头，那么需要反转链表
// 3.要反转链表的话，需要先拿到右半段的链表头，所以用快慢指针

func main() {
	a1, a2, a3, a4 := new(ListNode0234), new(ListNode0234), new(ListNode0234), new(ListNode0234)
	a1.Next = a2
	a2.Next = a3
	a3.Next = a4
	a1.Val, a2.Val, a3.Val, a4.Val = 1, 2, 2, 1
	for p := a1; p != nil; p = p.Next {
		fmt.Print(p.Val, " ")
	}
	fmt.Println()
	fmt.Println(isPalindrome(a1))
}
