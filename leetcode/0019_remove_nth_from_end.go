package main

import "fmt"

type ListNode0019 struct {
	Val  int
	Next *ListNode0019
}

// 标准解法：双指针距离差法，通过虚拟链表头避免删除链表头结点的问题
func removeNthFromEnd(head *ListNode0019, n int) *ListNode0019 {
	dummy := &ListNode0019{Next: head}
	a, b := dummy, dummy
	for i := 0; i <= n; i++ {
		b = b.Next
	}
	for b != nil {
		a = a.Next
		b = b.Next
	}
	a.Next = a.Next.Next
	return dummy.Next
}

// 思路总结：一般涉及链表增删的时候，要考虑是否引入虚拟头结点
func main() {
	a1, a2, a3, a4, a5 := new(ListNode0019), new(ListNode0019), new(ListNode0019), new(ListNode0019), new(ListNode0019)
	a1.Val, a2.Val, a3.Val, a4.Val, a5.Val = 1, 2, 3, 4, 5
	a1.Next = a2
	a2.Next = a3
	a3.Next = a4
	a4.Next = a5
	for p := removeNthFromEnd(a1, 2); p != nil; p = p.Next {
		fmt.Print(p.Val, " ")
	}
}
