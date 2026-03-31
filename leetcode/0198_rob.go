package main

import "fmt"

// 标准解法：动态规划
func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return max(nums[0], nums[1])
	}
	p, q, r := nums[0], max(nums[0], nums[1]), 0
	for i := 2; i < n; i++ {
		if p+nums[i] > q {
			r = p + nums[i]
		} else {
			r = q
		}
		p = q
		q = r
	}
	return q
}

// 思路总结：
// 1.动态规划的当前状态取决于过去，所以思考问题的时候不要看未来，贪心才看未来
// 2.一般动态规划都能用滚动变量，将空间复杂度化为常数级

func main() {
	nums := []int{2, 7, 9, 3, 1}
	fmt.Println(rob(nums))
}
