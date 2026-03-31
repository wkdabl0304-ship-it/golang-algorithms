package main

import (
	"fmt"
)

// 标准解法：动态规划
func coinChange(coins []int, amount int) int {
	//   i   = previous + x
	// dp[i] =  dp[i-x] + 1
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
	}
	for i := 1; i <= amount; i++ {
		need := amount + 1
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				x := i - coins[j]
				if dp[x]+1 < need {
					need = dp[x] + 1
				}
			}
		}
		dp[i] = need
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}

// 思路总结：难点主要是如何实现“凑不成”，其实就是初始化的时候用“无穷大”占位

func main() {
	coins := []int{1, 2, 5}
	amount := 11
	fmt.Println(coinChange(coins, amount))
}
