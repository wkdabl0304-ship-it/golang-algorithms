package main

import (
	"fmt"
)

// 心得体会
// 1.链表交叉问题可以用双指针相遇的特殊方法
// 2.用 for 遍历链表的时候，判断条件用当前指针，不要用下一个节点，不然最后还得补判断

type ListNode struct {
	Val  int
	Next *ListNode
}

// 1.个人解法(Hash法)
//func getIntersectionNode(headA, headB *ListNode) *ListNode {
//	m := make(map[*ListNode]bool)
//	for headB.Next != nil {
//		m[headB] = true
//		headB = headB.Next
//	}
//	m[headB] = true
//	for headA.Next != nil {
//		if m[headA] {
//			break
//		}
//		headA = headA.Next
//	}
//	if m[headA] {
//		return headA
//	} else {
//		return nil
//	}
//}

// 2.个人解法优化版
//func getIntersectionNode(headA, headB *ListNode) *ListNode {
//	m := make(map[*ListNode]bool)
//	for cur := headB; cur != nil; cur = cur.Next {
//		m[cur] = true
//	}
//	for cur := headA; cur != nil; cur = cur.Next {
//		if m[cur] {
//			return cur
//		}
//	}
//	return nil
//}

// 3.标准解法(双指针法)
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pA, pB := headA, headB
	for pA != pB {
		if pA == nil {
			pA = headB
		} else {
			pA = pA.Next
		}
		if pB == nil {
			pB = headA
		} else {
			pB = pB.Next
		}
	}
	return pA
}

func main() {
	a1, a2, b1, b2, b3, c1, c2, c3 := new(ListNode), new(ListNode), new(ListNode), new(ListNode), new(ListNode), new(ListNode), new(ListNode), new(ListNode)
	a1.Next, a2.Next, c1.Next, c2.Next, b1.Next, b2.Next, b3.Next = a2, c1, c2, c3, b2, b3, c1
	res := getIntersectionNode(a1, b1)
	fmt.Println(res)
}
