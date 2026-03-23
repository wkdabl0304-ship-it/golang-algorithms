package main

import "fmt"

// 标准解法：实现一个能够记录 map 里结点的时间顺序的队列，因为要反复移动，所以选择双向链表

type ListNode0146 struct {
	Key      int
	Val      int
	Next     *ListNode0146
	Previous *ListNode0146
}

type LRUCache struct {
	m        map[int]*ListNode0146
	head     *ListNode0146
	tail     *ListNode0146
	capacity int
	total    int
}

// Constructor0146 LRUCache 创建器
func Constructor0146(capacity int) LRUCache {
	dummyHead := new(ListNode0146)
	dummyTail := new(ListNode0146)
	dummyHead.Next = dummyTail
	dummyTail.Previous = dummyHead
	return LRUCache{
		m:        make(map[int]*ListNode0146, capacity),
		head:     dummyHead,
		tail:     dummyTail,
		capacity: capacity,
		total:    0,
	}
}

func (c *LRUCache) Get(key int) int {
	if v, ok := c.m[key]; ok {
		// 移动至链表尾
		v.Previous.Next = v.Next
		v.Next.Previous = v.Previous
		tailPrevious := c.tail.Previous
		tailPrevious.Next = v
		v.Previous = tailPrevious
		v.Next = c.tail
		c.tail.Previous = v
		return v.Val
	}
	return -1
}

func (c *LRUCache) Put(key int, value int) {
	if v := c.Get(key); v != -1 {
		// 键已存在，更新
		c.m[key].Val = value
	} else {
		// 键不存在，新增
		node := new(ListNode0146)
		node.Key, node.Val = key, value
		tailPrevious := c.tail.Previous
		tailPrevious.Next = node
		node.Previous = tailPrevious
		node.Next = c.tail
		c.tail.Previous = node
		c.m[key] = node
		c.total++
		// 2.去旧
		if c.total > c.capacity {
			headNext := c.head.Next
			headNextNext := c.head.Next.Next
			c.head.Next = headNextNext
			headNextNext.Previous = c.head
			headNext.Next, headNext.Previous = nil, nil
			delete(c.m, headNext.Key)
			c.total--
		}
	}
}

func main() {
	obj := Constructor0146(3)
	obj.Put(666, 233)
	param1 := obj.Get(666)
	fmt.Println(param1)
}
