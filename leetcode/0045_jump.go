package main

import "fmt"

// 标准解法：贪心算法
func jump(nums []int) int {
	n := len(nums)
	res, end := 0, 0
	maxReach := nums[0]
	for i := 0; i < n-1; i++ {
		if i+nums[i] > maxReach {
			maxReach = i + nums[i]
		}
		if i == end {
			end = maxReach
			res++
		}
	}
	return res
}

// 思路总结：
// 1.跳跃的时机不是更新 maxReach 的时候，而是必须跳的时候
// 2.maxReach 指的是所有可选步骤里能跳最远的位置

func main() {
	nums := []int{2, 3, 1, 1, 4}
	//nums := []int{1, 2}
	fmt.Println(jump(nums))
}
