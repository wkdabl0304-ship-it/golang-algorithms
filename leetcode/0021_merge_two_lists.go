package main

import "fmt"

type ListNode0021 struct {
	Val  int
	Next *ListNode0021
}

// 标准解法：双指针
func mergeTwoLists(list1 *ListNode0021, list2 *ListNode0021) *ListNode0021 {
	dummy := new(ListNode0021)
	current := dummy
	a, b := list1, list2
	for a != nil && b != nil {
		if b.Val <= a.Val {
			current.Next = b
			b = b.Next
		} else {
			current.Next = a
			a = a.Next
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

func main() {
	a1, a2, a3, b1, b2, b3 := new(ListNode0021), new(ListNode0021), new(ListNode0021), new(ListNode0021), new(ListNode0021), new(ListNode0021)
	a1.Val, a2.Val, a3.Val, b1.Val, b2.Val, b3.Val = 1, 2, 4, 1, 3, 4
	a1.Next = a2
	a2.Next = a3
	b1.Next = b2
	b2.Next = b3
	for p := mergeTwoLists(a1, b1); p != nil; p = p.Next {
		fmt.Print(p.Val, " ")
	}
}
