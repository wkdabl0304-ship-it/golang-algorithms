package main

import (
	"container/heap"
	"fmt"
)

type ListNode0023 struct {
	Val  int
	Next *ListNode0023
}

// 定义小顶堆
type minHeap0023 []*ListNode0023

func (h *minHeap0023) Len() int {
	return len(*h)
}
func (h *minHeap0023) Less(i, j int) bool {
	return (*h)[i].Val < (*h)[j].Val
}
func (h *minHeap0023) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}
func (h *minHeap0023) Push(x interface{}) {
	*h = append(*h, x.(*ListNode0023))
}
func (h *minHeap0023) Pop() interface{} {
	old := *h
	*h = old[0 : len(old)-1]
	return old[len(old)-1]
}

// 标准解法：结点池 + 归并排序
func mergeKLists(lists []*ListNode0023) *ListNode0023 {
	// 1.初始化结点池
	h := &minHeap0023{}
	heap.Init(h)
	for i := 0; i < len(lists); i++ {
		if lists[i] != nil {
			heap.Push(h, lists[i])
		}
	}
	// 2.串联结点
	dummy := new(ListNode0023)
	current := dummy
	for h.Len() > 0 {
		current.Next = heap.Pop(h).(*ListNode0023)
		if current.Next.Next != nil {
			heap.Push(h, current.Next.Next)
		}
		current = current.Next
	}
	return dummy.Next
}

// 思路总结：题目难点主要是如何在结点池里直接找到最小的结点，所以引入小顶堆

func main() {
	a1, a2, a3, b1, b2, b3, c1, c2 := new(ListNode0023), new(ListNode0023), new(ListNode0023), new(ListNode0023), new(ListNode0023), new(ListNode0023), new(ListNode0023), new(ListNode0023)
	a1.Val, a2.Val, a3.Val, b1.Val, b2.Val, b3.Val, c1.Val, c2.Val = 1, 4, 5, 1, 3, 4, 2, 6
	lists := []*ListNode0023{a1, b1, c1}
	a1.Next, a2.Next = a2, a3
	b1.Next, b2.Next = b2, b3
	c1.Next = c2
	for p := mergeKLists(lists); p != nil; p = p.Next {
		fmt.Print(p.Val, " ")
	}
}
