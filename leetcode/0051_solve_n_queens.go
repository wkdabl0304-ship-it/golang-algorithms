package main

import "fmt"

// 标准解法：回溯算法
func solveNQueens(n int) [][]string {
	var res [][]string
	queens := make([]int, n)
	column, diagonal1, diagonal2 := make(map[int]bool), make(map[int]bool), make(map[int]bool)
	var backtrack func(row int)
	backtrack = func(row int) {
		if row == n {
			oneCase := make([]string, 0, n)
			for i := 0; i < n; i++ {
				oneRow := make([]byte, 0, n)
				for j := 0; j < n; j++ {
					if j == queens[i] {
						oneRow = append(oneRow, 'Q')
					} else {
						oneRow = append(oneRow, '.')
					}
				}
				oneCase = append(oneCase, string(oneRow))
			}
			res = append(res, oneCase)
			return
		}
		for col := 0; col < n; col++ {
			if column[col] || diagonal1[row-col] || diagonal2[row+col] {
				continue
			}
			queens[row] = col
			column[col], diagonal1[row-col], diagonal2[row+col] = true, true, true
			backtrack(row + 1)
			column[col], diagonal1[row-col], diagonal2[row+col] = false, false, false
		}
	}
	backtrack(0)
	return res
}

// 思路总结：

func main() {
	n := 4
	fmt.Println(solveNQueens(n))
}
