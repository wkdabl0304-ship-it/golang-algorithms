package main

import "fmt"

func searchRange(nums []int, target int) []int {
	right := search0034(nums, target, true) - 1
	left := search0034(nums, target, false)
	if left <= right && left >= 0 && left < len(nums) && nums[left] == target {
		return []int{left, right}
	}
	return []int{-1, -1}
}

func search0034(nums []int, target int, isRight bool) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if target > nums[mid] || (isRight && target == nums[mid]) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

// 思路总结：二分查找如果将等于的情况并进大于或小于其中一种，相当于找左右边界了，注意一下边界就可以封装

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	fmt.Println(searchRange(nums, target))
}
