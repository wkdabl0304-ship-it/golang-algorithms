package main

import "fmt"

// 个人解法：右哨兵
//func largestRectangleArea(heights []int) int {
//	var stack []int
//	res := 0
//	newHeights := make([]int, len(heights)+1)
//	copy(newHeights, heights)
//	for i, v := range newHeights {
//		for len(stack) > 0 && v < newHeights[stack[len(stack)-1]] {
//			right := i - 1
//			current := stack[len(stack)-1]
//			left := 0
//			if len(stack) > 1 {
//				left = stack[len(stack)-2] + 1
//			}
//			area := (right - left + 1) * newHeights[current]
//			if area > res {
//				res = area
//			}
//			stack = stack[:len(stack)-1]
//		}
//		stack = append(stack, i)
//	}
//	return res
//}

// 标准解法: 双哨兵
func largestRectangleArea(heights []int) int {
	var stack []int
	res := 0
	newHeights := make([]int, len(heights)+2)
	copy(newHeights[1:], heights)
	for right, v := range newHeights {
		for len(stack) > 0 && v < newHeights[stack[len(stack)-1]] {
			current := stack[len(stack)-1]
			left := stack[len(stack)-2] + 1
			area := (right - left) * newHeights[current]
			if area > res {
				res = area
			}
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, right)
	}
	return res
}

// 思路总结：
// 1.通过单调栈找到第一个比柱子矮的，那么出栈元素的左一位的就是右边界
// 2.出栈元素的栈内前一位的右一位就是左边界
// 3.如果持续递增将没办法算面积，所以引入右哨兵

func main() {
	heights := []int{2, 1, 5, 6, 2, 3}
	fmt.Println(largestRectangleArea(heights))
}
