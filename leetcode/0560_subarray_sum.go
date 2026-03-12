package main

import "fmt"

// 标准解法：前缀和思想
func subarraySum(nums []int, k int) int {
	res := 0
	m := make(map[int]int)
	m[0] = 1
	currentTotal := 0
	for i := 0; i < len(nums); i++ {
		currentTotal += nums[i]
		if v, ok := m[currentTotal-k]; ok {
			res += v
		}
		m[currentTotal]++
	}
	return res
}

// 解题思路：通过 map + 前缀和 去除嵌套双循环，

func main() {
	nums := []int{1, 2, 3}
	k := 3
	fmt.Println(subarraySum(nums, k))
}
