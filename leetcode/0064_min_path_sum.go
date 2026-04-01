package main

import "fmt"

// 标准解法：多维动态规划
func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	dp := make([]int, n)
	dp[0] = grid[0][0]
	for i := 1; i < n; i++ {
		dp[i] = dp[i-1] + grid[0][i]
	}
	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			if j == 0 {
				dp[j] = dp[j] + grid[i][0]
			} else {
				up := dp[j]
				left := dp[j-1]
				dp[j] = min(up, left) + grid[i][j]
			}
		}
	}
	return dp[n-1]
}

// 思路总结：这道题依然可以用滚动更新，只是需要对第 0 列多做处理

func main() {
	grid := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	fmt.Println(minPathSum(grid))
}
