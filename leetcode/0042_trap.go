package main

// 标准解法：双指针法
func trap(height []int) int {
	res := 0
	left, right := 0, len(height)-1
	maxLeftHeight, maxRightHeight := 0, 0
	for left < right {
		if height[left] > maxLeftHeight {
			maxLeftHeight = height[left]
		}
		if height[right] > maxRightHeight {
			maxRightHeight = height[right]
		}
		if height[left] < height[right] {
			res += maxLeftHeight - height[left]
			left++
		} else {
			res += maxRightHeight - height[right]
			right--
		}
	}
	return res
}

// 解题思路：
// 1.以每一格为单位来思考能装多少水
// 2.水高取决于左右挡板较短的那一根
// 3.通过高度贪心实现剪枝
// 4.通过代码的先后顺序巧妙避免了负数高度
