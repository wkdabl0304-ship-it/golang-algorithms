package main

import "fmt"

// 标准解法：摩尔投票算法
func majorityElement(nums []int) int {
	count, candidate := 0, 0
	for _, v := range nums {
		if count == 0 {
			candidate = v
		}
		if v != candidate {
			count--
		} else {
			count++
		}
	}
	return candidate
}

// 思路总结：异项相消，留下来的就是最多数的

func main() {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println(majorityElement(nums))
}
