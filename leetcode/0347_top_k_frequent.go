package main

import (
	"container/heap"
	"fmt"
)

// 标准解法：小顶堆
type minHeap0347 []elem0347

type elem0347 struct {
	num       int // 数字
	frequency int // 频率
}

func (h *minHeap0347) Len() int {
	return len(*h)
}

func (h *minHeap0347) Less(i, j int) bool {
	return (*h)[i].frequency < (*h)[j].frequency // 下标为 i 的元素是否要排在下标为 j 的元素前面
}

func (h *minHeap0347) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *minHeap0347) Push(x interface{}) {
	*h = append(*h, x.(elem0347))
}

func (h *minHeap0347) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	h := new(minHeap0347)
	heap.Init(h)
	m := make(map[int]int, len(nums))
	var res []int
	for _, v := range nums {
		m[v]++
	}
	for key, value := range m {
		e := elem0347{
			num:       key,
			frequency: value,
		}
		heap.Push(h, e)
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	for i := k - 1; i >= 0; i-- {
		temp := heap.Pop(h).(elem0347)
		res = append(res, temp.num)
	}
	return res
}

// 思路总结：
// 1.最好先初始化好数组再初始化堆
// 2.堆空间要限制，不然浪费空间

func main() {
	nums := []int{1, 2, 1, 2, 1, 2, 3, 1, 3, 2}
	k := 2
	fmt.Println(topKFrequent(nums, k))
}
