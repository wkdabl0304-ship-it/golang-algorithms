package main

import "fmt"

// 标准解法：动态规划
func canPartition(nums []int) bool {
	total := 0
	for _, v := range nums {
		total += v
	}
	if total%2 == 1 {
		return false
	}
	target := total / 2
	dp := make([]bool, target+1)
	dp[0] = true
	for _, v := range nums {
		for j := target; j >= v; j-- {
			if dp[j-v] {
				dp[j] = true
			}
		}
		if dp[target] {
			return true
		}
	}
	return false
}

// 思路总结：
// 1.本来应该用二维数组，但是背包问题只和上一次状态有关，所以可以滚动更新
// 2.之所以倒序遍历背包容量，是为了保证 dp[j-v] 是没有添加当前新物品的状态

func main() {
	nums := []int{1, 5, 11, 5}
	fmt.Println(canPartition(nums))
}
