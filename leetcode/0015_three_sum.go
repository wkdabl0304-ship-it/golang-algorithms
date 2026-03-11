package main

import (
	"fmt"
	"sort"
)

// 个人解法：问题是没有解决去重问题
//func threeSum(nums []int) [][]int {
//	m := make(map[int]int)
//	for i, v := range nums {
//		m[v] = i
//	}
//	var res [][]int
//	for slow := 0; slow < len(nums); slow++ {
//		for fast := slow + 1; fast < len(nums); fast++ {
//			if i, ok := m[-(nums[slow] + nums[fast])]; ok && i > fast {
//				res = append(res, []int{nums[slow], nums[fast], nums[i]})
//			}
//		}
//	}
//	return res
//}

// 标准解法：通过排序解决去重问题，根据 sum 的大小判断移动左右指针（贪心剪枝）
func threeSum(nums []int) [][]int {
	// 1.排序
	var res [][]int
	sort.Ints(nums)
	// 2.固定一个指针，移动左指针和右指针
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				for nums[left] == nums[left+1] && left < right {
					left++
				}
				for nums[right] == nums[right-1] && left < right {
					right--
				}
				left++
				right--
			} else if sum > 0 {
				right--
			} else {
				left++
			}
		}
	}
	return res
}

// 做法总结：双指针问题一般可以通过已知解来贪心剪枝

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}
