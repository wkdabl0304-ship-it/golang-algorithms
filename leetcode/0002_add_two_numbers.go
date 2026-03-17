package main

import "fmt"

type ListNode0002 struct {
	Val  int
	Next *ListNode0002
}

// 标准解法：创建新链表法，原地修改法要处理多种拓扑情况
func addTwoNumbers(l1 *ListNode0002, l2 *ListNode0002) *ListNode0002 {
	a, b := l1, l2
	carry := 0
	dummy := new(ListNode0002)
	current := dummy
	for a != nil || b != nil || carry != 0 {
		sum := carry
		if a != nil {
			sum += a.Val
			a = a.Next
		}
		if b != nil {
			sum += b.Val
			b = b.Next
		}
		current.Next = &ListNode0002{Val: sum % 10}
		carry = sum / 10
		current = current.Next
	}
	return dummy.Next
}

// 思路总结：当题目没有要求原地修改时，没必要去原地修改

func main() {
	a1, a2, a3, b1, b2, b3 := new(ListNode0002), new(ListNode0002), new(ListNode0002), new(ListNode0002), new(ListNode0002), new(ListNode0002)
	a1.Val, a2.Val, a3.Val, b1.Val, b2.Val, b3.Val = 2, 4, 3, 5, 6, 4
	a1.Next = a2
	a2.Next = a3
	b1.Next = b2
	b2.Next = b3
	for p := addTwoNumbers(a1, b1); p != nil; p = p.Next {
		fmt.Print(p.Val, " ")
	}
}
