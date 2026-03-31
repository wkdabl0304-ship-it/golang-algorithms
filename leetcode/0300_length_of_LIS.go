package main

import "fmt"

// 标准解法：动态规划
func lengthOfLIS(nums []int) int {
	// dp[i] = (dp[i-j] + 1) or 1
	n := len(nums)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	maxLen := 1
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
			}
		}
		if dp[i] > maxLen {
			maxLen = dp[i]
		}
	}
	return maxLen
}

// 思路总结：
// 1.易错：dp[i] 不只是和前一种情况挂钩，而是所有情况
// 2.dp[i] 随时可能断开

func main() {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Println(lengthOfLIS(nums))
}
