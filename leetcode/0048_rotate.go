package main

import "fmt"

// 标准解法：几何拆解法
func rotate0048(matrix [][]int) {
	n := len(matrix)
	// 1.上下翻转
	for i := 0; i <= (n-1)/2; i++ {
		for j := 0; j < n; j++ {
			matrix[i][j], matrix[n-1-i][j] = matrix[n-1-i][j], matrix[i][j]
		}
	}
	// 2.主对角线翻转
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate0048(matrix)
	fmt.Println(matrix)
}
