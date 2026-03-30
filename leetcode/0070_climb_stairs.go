package main

import "fmt"

// 标准解法：动态规划
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	p, q, r := 1, 2, 0
	for i := 3; i <= n; i++ {
		r = p + q
		p = q
		q = r
	}
	return q
}

// 思路总结：
// 1.动态规划的基本特征就是，当前的结果取决于前者的结果
// 2.动态规划包括了贪心，动态规划一般解决“结果数量”问题，贪心解决“最”问题

func main() {
	n := 3
	fmt.Println(climbStairs(n))
}
