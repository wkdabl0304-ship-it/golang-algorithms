package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 标准解法：快慢指针追逐法，避免了用 map
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

func main() {
	a1, a2, a3, a4 := new(ListNode), new(ListNode), new(ListNode), new(ListNode)
	a1.Val, a2.Val, a3.Val, a4.Val = 3, 2, 0, -4
	a1.Next = a2
	a2.Next = a3
	a3.Next = a4
	a4.Next = a2
	fmt.Println(hasCycle(a1))
}
