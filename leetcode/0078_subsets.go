package main

import "fmt"

// 标准解法：回溯算法
func subsets(nums []int) [][]int {
	var res [][]int
	path := make([]int, 0, len(nums))
	var backtrack func(start int)
	backtrack = func(start int) {
		temp := make([]int, len(path))
		copy(temp, path)
		res = append(res, temp)
		for i := start; i < len(nums); i++ {
			path = append(path, nums[i])
			backtrack(i + 1)
			path = path[:len(path)-1]
		}
	}
	backtrack(0)
	return res
}

// 思路总结：组合不需要标记，只有排列才需要标记

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(subsets(nums))
}
