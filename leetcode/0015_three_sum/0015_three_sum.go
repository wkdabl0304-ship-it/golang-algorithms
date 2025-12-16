package main

import (
	"fmt"
	"sort"
)

// 心得体会
// 1.解题先想暴力解法，再想怎么优化
// 2.先写出主流程，再慢慢补充边界条件
// 3.有条件性地去重一般用排序+指针

func threeSum(nums []int) [][]int {

	var res [][]int
	sort.Ints(nums)
	length := len(nums)
	for i := 0; i < length-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, length-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return res

}

func main() {
	//nums := []int{-1, 0, 1, 2, -1, -4}
	nums := []int{0, 0, 0}
	res := threeSum(nums)
	fmt.Println(res)
}
