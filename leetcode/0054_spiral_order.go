package main

import "fmt"

// 标准解法：动态边界法
func spiralOrder(matrix [][]int) []int {
	res := make([]int, 0, len(matrix)*len(matrix[0]))
	up, down, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1
	for {
		// 1.向右移
		for j := left; j <= right; j++ {
			res = append(res, matrix[up][j])
		}
		up++
		if up > down {
			break
		}
		// 2.向下移
		for i := up; i <= down; i++ {
			res = append(res, matrix[i][right])
		}
		right--
		if right < left {
			break
		}
		// 3. 向左移
		for j := right; j >= left; j-- {
			res = append(res, matrix[down][j])
		}
		down--
		if down < up {
			break
		}
		// 4. 向上移
		for i := down; i >= up; i-- {
			res = append(res, matrix[i][left])
		}
		left++
		if left > right {
			break
		}
	}
	return res
}

func main() {
	matrix := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	fmt.Println(spiralOrder(matrix))
}
