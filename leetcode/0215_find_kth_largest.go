package main

import (
	"container/heap"
	"fmt"
)

// 最优解法：快排对半选择法
//func findKthLargest(nums []int, k int) int {
//	target := len(nums) - k // target 指的是第 k 大的数字在排序后的下标
//	return quickSort(nums, 0, len(nums)-1, target)
//}
//
//func quickSort(nums []int, left, right, target int) int {
//	if left == right {
//		return nums[left]
//	}
//	p := partition(nums, left, right)
//	if p == target {
//		return nums[p]
//	} else if p < target {
//		return quickSort(nums, p+1, right, target)
//	} else {
//		return quickSort(nums, left, p-1, target)
//	}
//}
//
//func partition(nums []int, left, right int) int {
//	pivotIdx := left + rand.Intn(right-left+1)
//	pivot := nums[pivotIdx]
//	nums[pivotIdx], nums[right] = nums[right], nums[pivotIdx]
//	i := left
//	for j := left; j < right; j++ {
//		if nums[j] < pivot {
//			nums[i], nums[j] = nums[j], nums[i]
//			i++
//		}
//	}
//	nums[i], nums[right] = nums[right], nums[i]
//	return i
//}

// 面试解法：小顶堆
type minHeap0215 []int

func (h *minHeap0215) Len() int {
	return len(*h)
}

func (h *minHeap0215) Less(i, j int) bool {
	return (*h)[i] < (*h)[j] // 下标为 i 的元素是否要排在下标为 j 的元素前面
}

func (h *minHeap0215) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *minHeap0215) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *minHeap0215) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

func findKthLargest(nums []int, k int) int {
	h := new(minHeap0215)
	heap.Init(h)
	for _, v := range nums {
		heap.Push(h, v)
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	return (*h)[0]
}

// 思路总结：
// 1.最优方法使用二分选择法，但是边界以及范围较难控制
// 2.工程最优方法应该用小顶堆

func main() {
	nums := []int{3, 2, 1, 5, 6, 4}
	k := 2
	fmt.Println(findKthLargest(nums, k))
}
