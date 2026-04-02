package main

import "fmt"

// 标准解法：多维动态规划
func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m][n]
}

// 思路总结：
// 1.首先要明白如何手动比对两个字符串，就是地毯式地一个一个比
// 2.接着就会发现计数的时候后者取决于前者，所以用 DP
// 3.因为要比对两个字符串，很自然地就想到用 dp[i][j] 了

func main() {
	text1 := "abcde"
	text2 := "ace"
	fmt.Println(longestCommonSubsequence(text1, text2))
}
