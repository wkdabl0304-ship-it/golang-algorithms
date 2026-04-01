package main

import "fmt"

// 标准解法：多维动态规划
func uniquePaths(m int, n int) int {
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[j] += dp[j-1]
		}
	}
	return dp[n-1]
}

// 思路总结：当前的路径数取决于左边和上方的格子的数量之和

func main() {
	m := 3
	n := 7
	fmt.Println(uniquePaths(m, n))
}
