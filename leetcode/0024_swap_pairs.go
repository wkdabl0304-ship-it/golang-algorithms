package main

import "fmt"

type ListNode0024 struct {
	Val  int
	Next *ListNode0024
}

// 标准解法：三指针迭代法
func swapPairs(head *ListNode0024) *ListNode0024 {
	dummy := &ListNode0024{Next: head}
	previous := dummy
	for previous.Next != nil && previous.Next.Next != nil {
		// 1.初始化
		current := previous.Next
		next := current.Next
		// 2.交换
		previous.Next = next
		current.Next = next.Next
		next.Next = current
		// 3.更新指针
		previous = current
	}
	return dummy.Next
}

// 思路总结：出现多指针时，为了避免追踪混乱，最好锚定一个基准指针
func main() {
	a1, a2, a3, a4 := new(ListNode0024), new(ListNode0024), new(ListNode0024), new(ListNode0024)
	a1.Val, a2.Val, a3.Val, a4.Val = 1, 2, 3, 4
	a1.Next = a2
	a2.Next = a3
	a3.Next = a4
	for p := swapPairs(a2); p != nil; p = p.Next {
		fmt.Print(p.Val, " ")
	}
}
