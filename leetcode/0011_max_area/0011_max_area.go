package main

import "fmt"

// 心得体会
// 1.把穷举变成一步一步找，能排除掉一些绝对错误的选项

func maxArea(height []int) int {

	//1.穷举法，时间复杂度O(n2)
	//maxA := 0
	//for i := 0; i < len(height)-1; i++ {
	//	for j := i + 1; j < len(height); j++ {
	//		minHeight := height[i]
	//		if height[i] > height[j] {
	//			minHeight = height[j]
	//		}
	//		area := minHeight * (j - i)
	//		if area > maxA {
	//			maxA = area
	//		}
	//	}
	//}
	//return maxA

	//2.安全贪心算法：优化掉了短板移动后更矮但计算面积的过程
	//i, j := 0, len(height)-1
	//maxA := 0
	//for i < j {
	//	minH := height[i]
	//	if height[i] > height[j] {
	//		minH = height[j]
	//	}
	//	area := minH * (j - i)
	//	if area > maxA {
	//		maxA = area
	//	}
	//	if height[i] < height[j] {
	//		for i < j {
	//			prevH := height[i]
	//			i++
	//			if height[i] > prevH {
	//				break
	//			}
	//		}
	//	} else {
	//		for i < j {
	//			prevH := height[j]
	//			j--
	//			if height[j] > prevH {
	//				break
	//			}
	//		}
	//	}
	//}
	//return maxA

	//3.标准解法
	i, j := 0, len(height)-1
	maxA := 0

	for i < j {
		h := height[i]
		if height[j] < h {
			h = height[j]
		}
		area := h * (j - i)
		if area > maxA {
			maxA = area
		}

		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}

	return maxA
}

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	res := maxArea(height)
	fmt.Println(res)
}
