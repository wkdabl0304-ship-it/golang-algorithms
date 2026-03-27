package main

import (
	"fmt"
	"math/rand"
)

// 标准解法：快排对半选择法
func findKthLargest(nums []int, k int) int {
	target := len(nums) - k // target 指的是第 k 大的数字在排序后的下标
	return quickSort(nums, 0, len(nums)-1, target)
}

func quickSort(nums []int, left, right, target int) int {
	if left == right {
		return nums[left]
	}
	p := partition(nums, left, right)
	if p == target {
		return nums[p]
	} else if p < target {
		return quickSort(nums, p+1, right, target)
	} else {
		return quickSort(nums, left, p-1, target)
	}
}

func partition(nums []int, left, right int) int {
	pivotIdx := left + rand.Intn(right-left+1)
	pivot := nums[pivotIdx]
	nums[pivotIdx], nums[right] = nums[right], nums[pivotIdx]
	i := left
	for j := left; j < right; j++ {
		if nums[j] < pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[i], nums[right] = nums[right], nums[i]
	return i
}

// 思路总结：这道题难就难在对快排边界的精确理解，而且得用单向快排比较好理解

func main() {
	nums := []int{3, 2, 1, 5, 6, 4}
	k := 2
	fmt.Println(findKthLargest(nums, k))
}
