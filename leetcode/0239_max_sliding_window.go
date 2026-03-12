package main

import "fmt"

// 标准解法：单调队列法
func maxSlidingWindow(nums []int, k int) []int {
	res := make([]int, 0, len(nums)-k+1)
	queue := make([]int, 0, k)
	for i, v := range nums {
		// 1.踢人
		for len(queue) > 0 && v >= nums[queue[len(queue)-1]] {
			queue = queue[:len(queue)-1]
		}
		// 2.入队
		queue = append(queue, i)
		// 3.去旧
		if queue[0] <= i-k {
			queue = queue[1:]
		}
		// 4.取最大
		if i+1 >= k {
			res = append(res, nums[queue[0]])
		}
	}
	return res
}

// 解题思路：
// 1.通过 队列 解决嵌套循环问题
// 2.通过 单调 实现比较数值并取最大
// 3.通过 递减 实现暂存滑动窗口的内容

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	fmt.Println(maxSlidingWindow(nums, k))
}
