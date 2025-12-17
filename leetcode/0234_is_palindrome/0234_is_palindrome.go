package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 0.复制数组法：时间复杂度O(n) 空间复杂度(n)

// 1.快慢指针法(栈版本)：时间复杂度O(n) 空间复杂度(n)
//func isPalindrome(head *ListNode) bool {
//	var stack []int
//	slow, fast := head, head
//	for fast != nil && fast.Next != nil {
//		stack = append(stack, slow.Val)
//		slow = slow.Next
//		fast = fast.Next.Next
//	}
//	if fast != nil {
//		slow = slow.Next
//	}
//	for len(stack) > 0 && slow.Val == stack[len(stack)-1] {
//		slow = slow.Next
//		stack = stack[:len(stack)-1]
//	}
//	return len(stack) == 0
//}

// 2.快慢指针(反转链表法+恢复反转)：时间复杂度O(n) 空间复杂度O(1)
//func isPalindrome(head *ListNode) bool {
//	slow, fast := head, head
//	for fast != nil && fast.Next != nil {
//		slow = slow.Next
//		fast = fast.Next.Next
//	}
//	if fast != nil {
//		slow = slow.Next
//	}
//	var prev *ListNode
//	for cur := slow; cur != nil; {
//		next := cur.Next
//		cur.Next = prev
//		prev = cur
//		cur = next
//	}
//	for start, end := head, prev; end != nil; {
//		if start.Val != end.Val {
//			return false
//		}
//		start = start.Next
//		end = end.Next
//	}
//	cur := prev
//	prev = nil
//	for cur != nil {
//		next := cur.Next
//		cur.Next = prev
//		prev = cur
//		cur = next
//	}
//	return true
//}

// 3.递归法
var left *ListNode

func dfs(right *ListNode) bool {
	if right == nil {
		return true
	}
	if !dfs(right.Next) {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	left = left.Next
	return true
}
func isPalindrome(head *ListNode) bool {
	left = head
	return dfs(head)
}

func main() {
	a4 := &ListNode{
		Val:  1,
		Next: nil,
	}
	a3 := &ListNode{
		Val:  2,
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
	res := isPalindrome(a1)
	fmt.Println(res)
}
