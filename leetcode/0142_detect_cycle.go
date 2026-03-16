package main

import "fmt"

type ListNode0142 struct {
	Val  int
	Next *ListNode0142
}

// 标准解法：快慢指针 + 追及问题，避免使用 map
func detectCycle(head *ListNode0142) *ListNode0142 {
	if head == nil {
		return nil
	}
	// 1.找到相遇点
	slow, fast, meet := head, head, (*ListNode0142)(nil)
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if slow == fast {
			meet = slow
			break
		}
	}
	if meet == nil {
		return nil
	}
	// 2.找到入口点
	a, b := head, meet
	for a != b {
		a, b = a.Next, b.Next
	}
	return a
}

// 思路总结
// 1.首先依据 LC0141 先找到相遇点
// 2.设 a 为起点到入口距离，b 为入口到相遇点距离，c 为相遇点到入口距离
// 3.那么有 s1 = a + b , s2 = a + n(b + c) + b
// 4.因为 2(a + b) = a + n(b + c) + b，可得 a = (n - 1)(b + c) + c
// 5.说明如果在环上走了 a 的距离和走 c 的距离，最终的落点一致
// 6.所以如果在相遇点走 a 距离，那么刚好走到入口，而从起点走 a 距离，也刚好到入口，刚好可以相遇

func main() {
	a1, a2, a3, a4 := new(ListNode0142), new(ListNode0142), new(ListNode0142), new(ListNode0142)
	a1.Val, a2.Val, a3.Val, a4.Val = 3, 2, 0, -4
	a1.Next = a2
	a2.Next = a3
	a3.Next = a4
	a4.Next = a2
	fmt.Println(detectCycle(a1))
}
