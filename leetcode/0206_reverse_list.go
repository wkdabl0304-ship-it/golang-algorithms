package main

import "fmt"

type ListNode0206 struct {
	Val  int
	Next *ListNode0206
}

// 标准解法：三指针迭代法
func reverseList(head *ListNode0206) *ListNode0206 {
	if head == nil {
		return nil
	}
	previous, current, next := (*ListNode0206)(nil), head, head.Next
	for current != nil {
		next = current.Next
		current.Next = previous
		previous = current
		current = next
	}
	return previous
}

func main() {
	a1, a2, a3, a4, a5 := new(ListNode0206), new(ListNode0206), new(ListNode0206), new(ListNode0206), new(ListNode0206)
	a1.Next = a2
	a2.Next = a3
	a3.Next = a4
	a4.Next = a5
	a1.Val, a2.Val, a3.Val, a4.Val, a5.Val = 1, 2, 3, 4, 5
	for p := a1; p != nil; p = p.Next {
		fmt.Print(p.Val, " ")
	}
	fmt.Println()
	for p := reverseList(a1); p != nil; p = p.Next {
		fmt.Print(p.Val, " ")
	}
}
