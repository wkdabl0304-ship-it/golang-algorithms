package main

import "fmt"

// 标准解法：动态规划
func numSquares(n int) int {
	// total = previous + x2(取最大，同时 x2<i )
	// dp[i] = dp[i-x2] + 1
	dp := make([]int, n+1)
	dp[1] = 1
	for i := 2; i <= n; i++ {
		num := dp[i-1] + 1
		for j := 1; j*j <= i; j++ {
			previous := i - j*j
			if dp[previous]+1 < num {
				num = dp[previous] + 1
			}
		}
		dp[i] = num
	}
	return dp[n]
}

// 思路总结：
// 1.做动态规划题的目标是找到状态转移方程
// 2.一般以目标结果作为 dp[i]

func main() {
	n := 12
	fmt.Println(numSquares(n))
}
