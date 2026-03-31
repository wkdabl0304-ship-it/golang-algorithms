package main

import "fmt"

// 标准解法：双dp
func maxProduct(nums []int) int {
	res := nums[0]
	currentMax, currentMin := res, res
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			currentMax, currentMin = currentMin, currentMax
		}
		currentMax = max(nums[i], nums[i]*currentMax)
		currentMin = min(nums[i], nums[i]*currentMin)
		if currentMax > res {
			res = currentMax
		}
	}
	return res
}

// 思路总结：相较于正常的dp，这道题需要双dp，因为随时切换

func main() {
	nums := []int{2, 3, -2, 4}
	fmt.Println(maxProduct(nums))
}
