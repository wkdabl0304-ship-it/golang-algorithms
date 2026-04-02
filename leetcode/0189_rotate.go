package main

import "fmt"

// 标准解法：三次反转法
func reverse0189(nums []int, left, right int) {
	i, j := left, right
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

func rotate0189(nums []int, k int) {
	t := k % len(nums)
	reverse0189(nums, 0, len(nums)-1)
	reverse0189(nums, 0, t-1)
	reverse0189(nums, t, len(nums)-1)
}

// 思路总结：遇到这种轮转的、截断拼接的，一般用反转

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	rotate0189(nums, k)
	fmt.Println(nums)
}
