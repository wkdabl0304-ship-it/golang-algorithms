package main

import "fmt"

// 心得体会
// 1.写递归的个人总结：1.写对自身函数的调用(传入下一代的值) 2.思考返回值与结束条件(构成递归) 3.完善递归逻辑
// 2.迭代的时候因为指针过多容易混乱，应当以一个指针(如 cur)作为核心来组织思路

type ListNode struct {
	Val  int
	Next *ListNode
}

// 1.递归法(通俗易懂版)
//func reverseList(head *ListNode) *ListNode {
//	if head == nil || head.Next == nil {
//		return head
//	}
//	p := head
//	q := p.Next
//	newH := reverseList(q)
//	q.Next = p
//	p.Next = nil
//	return newH
//}

// 2.递归法(标准版，内存消耗更少)
//func reverseList(head *ListNode) *ListNode {
//	if head == nil || head.Next == nil {
//		return head
//	}
//	newH := reverseList(head.Next)
//	head.Next.Next = head
//	head.Next = nil
//	return newH
//}

// 3.迭代法(个人版)
//func reverseList(head *ListNode) *ListNode {
//	if head == nil || head.Next == nil {
//		return head
//	}
//	p := head
//	q := p.Next
//	p.Next = nil
//	for q != nil {
//		t := q.Next
//		q.Next = p
//		p = q
//		q = t
//	}
//	return p
//}

// 4.迭代法(标准版)
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}

func main() {
	a5 := &ListNode{
		Val:  5,
		Next: nil,
	}
	a4 := &ListNode{
		Val:  4,
		Next: a5,
	}
	a3 := &ListNode{
		Val:  3,
		Next: a4,
	}
	a2 := &ListNode{
		Val:  2,
		Next: a3,
	}
	a1 := &ListNode{
		Val:  1,
		Next: a2,
	}
	res := reverseList(a1)
	fmt.Println(res)
}
