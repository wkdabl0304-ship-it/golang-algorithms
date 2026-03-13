package main

import "fmt"

// 标准解法：原地标记法
func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	row0, column0 := false, false
	// 1.处理标记行和标记列
	for j := 0; j < n; j++ {
		if matrix[0][j] == 0 {
			row0 = true
		}
	}
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			column0 = true
		}
	}
	// 2.标记
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	// 3.置零
	for j := 1; j < n; j++ {
		if matrix[0][j] == 0 {
			for t := 1; t < m; t++ {
				matrix[t][j] = 0
			}
		}
	}
	for i := 1; i < m; i++ {
		if matrix[i][0] == 0 {
			for t := 1; t < n; t++ {
				matrix[i][t] = 0
			}
		}
	}
	// 4.处理标记行和标记列
	if row0 == true {
		for j := 0; j < n; j++ {
			matrix[0][j] = 0
		}
	}
	if column0 == true {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
}

// 思路总结：
// 0.首先想到的是直接改，但是直接改后的零会污染后面的清零步骤
// 1.后面想到的是矩阵复制，空间复杂度 O(mn)
// 2.接着可以通过只记录要标记的行和列，空间复杂度优化为 O(m+n)
// 3.因为一旦某个元素的行和列清零，实际上所有的数据不存在了，所以可以直接在第一行/第一列标记
// 但是会出现污染问题，也就是不知道 matrix[0][0] 上的 0 到底指的是行清零还是列清零
// 所以引入 row0 和 column0 来记录是哪个清零

func main() {
	matrix := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	setZeroes(matrix)
	fmt.Println(matrix)
}
