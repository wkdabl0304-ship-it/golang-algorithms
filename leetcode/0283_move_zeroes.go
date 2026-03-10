package main

import "fmt"

// 标准解法：交换法
func moveZeroes(nums []int) {
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
	}
}

// 解题思路：slow 指的是“下一个非零数“待的位置，而不是需要指向零
// 双指针核心逻辑：固定一个指针，动另外一个指针，如果一起动的话会变复杂

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)
}
