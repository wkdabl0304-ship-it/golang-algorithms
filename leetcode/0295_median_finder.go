package main

import (
	"container/heap"
	"fmt"
)

// 标准解法：对顶堆（两个有序堆），实现插入操作的时间复杂度为 O(logn)

type minHeap0295 []int

func (h *minHeap0295) Len() int {
	return len(*h)
}
func (h *minHeap0295) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}
func (h *minHeap0295) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}
func (h *minHeap0295) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *minHeap0295) Pop() interface{} {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}

type maxHeap0295 []int

func (h *maxHeap0295) Len() int {
	return len(*h)
}
func (h *maxHeap0295) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}
func (h *maxHeap0295) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}
func (h *maxHeap0295) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *maxHeap0295) Pop() interface{} {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}

type MedianFinder struct {
	maxH *maxHeap0295
	minH *minHeap0295
}

func Constructor0295() MedianFinder {
	maxH := &maxHeap0295{}
	minH := &minHeap0295{}
	heap.Init(maxH)
	heap.Init(minH)
	return MedianFinder{
		maxH: maxH,
		minH: minH,
	}
}

func (m *MedianFinder) AddNum(num int) {
	heap.Push(m.maxH, num)
	x := heap.Pop(m.maxH).(int)
	heap.Push(m.minH, x)
	if m.maxH.Len() < m.minH.Len() {
		y := heap.Pop(m.minH).(int)
		heap.Push(m.maxH, y)
	}
}

func (m *MedianFinder) FindMedian() float64 {
	if (m.maxH.Len()+m.minH.Len())%2 == 0 {
		return float64((*m.minH)[0]+(*m.maxH)[0]) / 2
	}
	return float64((*m.maxH)[0])
}

func main() {
	obj := Constructor0295()
	obj.AddNum(1)
	obj.AddNum(2)
	obj.AddNum(3)
	obj.AddNum(4)
	obj.AddNum(5)
	obj.AddNum(6)
	fmt.Println(obj.FindMedian())
}
