package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
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
	fmt.Println(searchMatrix(matrix, target))
}
