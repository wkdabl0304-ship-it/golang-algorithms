package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 标准解法：双指针法，通过相交链表的特殊性找到追及的共同点，避免了对 map 的使用
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	pA, pB := headA, headB
	for pA != pB {
		if pA != nil {
			pA = pA.Next
		} else {
			pA = headB
		}
		if pB != nil {
			pB = pB.Next
		} else {
			pB = headA
		}
	}
	return pA
}

func main() {
	a1, a2, c1, c2, c3 := new(ListNode), new(ListNode), new(ListNode), new(ListNode), new(ListNode)
	b1, b2 := new(ListNode), new(ListNode)
	a1.Next = a2
	a2.Next = c1
	c1.Next = c2
	c2.Next = c3
	b1.Next = b2
	b2.Next = c1
	fmt.Println(getIntersectionNode(a1, b1))
}
