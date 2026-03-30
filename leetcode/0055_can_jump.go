package main

import "fmt"

// 标准解法：贪心算法，回溯会导致时间复杂度太高
func canJump(nums []int) bool {
	maxReach := nums[0]
	n := len(nums)
	for i := 0; i < n && i <= maxReach; i++ {
		if i+nums[i] > maxReach {
			maxReach = i + nums[i]
		}
		if maxReach >= n-1 {
			return true
		}
	}
	return false
}

// 思路总结：如果每一步都选当前看起来最好的且不后悔，能够找到最优解，那么考虑用贪心

func main() {
	nums := []int{2, 3, 1, 1, 4}
	fmt.Println(canJump(nums))
}
