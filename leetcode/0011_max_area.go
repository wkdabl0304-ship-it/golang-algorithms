package main

// 标准解法：通过特殊的贪心来避免O(n2)的时间复杂度
func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxRes := 0
	for left < right {
		w := right - left
		h := 0
		if height[left] > height[right] {
			h = height[right]
			right--
		} else {
			h = height[left]
			left++
		}
		res := w * h
		if res > maxRes {
			maxRes = res
		}
	}
	return maxRes
}
