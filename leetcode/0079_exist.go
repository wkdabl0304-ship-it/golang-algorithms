package main

import "fmt"

// 标准解法：回溯算法
func exist(board [][]byte, word string) bool {
	rows, columns := len(board), len(board[0])
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	var backtrack func(i, j int, letter int) bool
	backtrack = func(i, j int, letter int) bool {
		if i < 0 || i >= rows || j < 0 || j >= columns || board[i][j] != word[letter] {
			return false
		}
		if letter == len(word)-1 {
			return true
		}
		temp := board[i][j]
		board[i][j] = '0'
		for _, v := range dirs {
			if backtrack(i+v[0], j+v[1], letter+1) {
				board[i][j] = temp
				return true
			}
		}
		board[i][j] = temp
		return false
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if board[i][j] == word[0] {
				if backtrack(i, j, 0) {
					return true
				}
			}
		}
	}
	return false
}

// 思路总结：
// 1.思考这种类型的题目的时候，要先处理当前结点，再递归往下
// 2.一般通过递归返回值可以去掉全局变量，而且剪枝提高效率

func main() {
	board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	word := "ABCCED"
	fmt.Println(exist(board, word))
}
