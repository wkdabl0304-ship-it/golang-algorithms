package main

import "fmt"

// 标准解法：阶梯法，通过线性搜索的方式避免了使用搜索算法，关键是要让选择有区分性
func searchMatrix0074(matrix [][]int, target int) bool {
	i, j := 0, len(matrix[0])-1
	for {
		if matrix[i][j] == target {
			return true
		} else if target > matrix[i][j] {
			i++
		} else {
			j--
		}
		if i == len(matrix) || j < 0 {
			break
		}
	}
	return false
}

func main() {
	matrix := [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}
	target := 5
	fmt.Println(searchMatrix0074(matrix, target))
}
