package main

import (
	"fmt"
)

// 个人解法：索引解耦
//func search(nums []int, target int) int {
//	n := len(nums)
//	k := searchStart(nums) // (坐标 - k + n) % n = 递增序坐标  坐标 = (递增序坐标 + k) % n
//	left, right := 0, n-1
//	for left <= right {
//		mid := left + (right-left)>>1
//		truthMid := (mid + k) % n
//		if target == nums[truthMid] {
//			return truthMid
//		} else if target > nums[truthMid] {
//			left = mid + 1
//		} else {
//			right = mid - 1
//		}
//	}
//	return -1
//}
//
//func searchStart(nums []int) int {
//	left, right := 0, len(nums)-1
//	for left < right {
//		mid := left + (right-left)>>1
//		if nums[mid] > nums[right] {
//			left = mid + 1
//		} else {
//			right = mid
//		}
//	}
//	return left
//}

// 标准解法：二分法
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if target == nums[mid] {
			return mid
		}
		if nums[left] <= nums[mid] {
			if target >= nums[left] && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

// 思路总结：巧妙思路：通过中点值来判断哪一侧是连续的，从而进行搜索

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0
	fmt.Println(search(nums, target))
}
