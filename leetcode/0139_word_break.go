package main

import "fmt"

// 标准解法：动态规划
func wordBreak(s string, wordDict []string) bool {
	// current = previous + x
	// dp[i] = dp[i-x] && x in wordDict
	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true
	for i := 1; i <= n; i++ {
		for j := 0; j < len(wordDict); j++ {
			start := i - len(wordDict[j])
			if start >= 0 && dp[start] && s[start:i] == wordDict[j] {
				dp[i] = true
				break
			}
		}
	}
	return dp[n]
}

// 思路总结：
// 1.易错：dp[i]的长度是 n+1，但是 s 的长度是 n，所以索引要区分是否加一减一
// 2.找到答案可以 break 进行优化

func main() {
	s := "applepenapple"
	wordDict := []string{"apple", "pen"}
	fmt.Println(wordBreak(s, wordDict))
}
