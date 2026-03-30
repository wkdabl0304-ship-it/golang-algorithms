package main

import "fmt"

// 标准解法：动态规划
func generate(numRows int) [][]int {
	var res [][]int
	for i := 0; i < numRows; i++ {
		row := make([]int, i+1)
		row[0], row[i] = 1, 1
		for j := 1; j < i; j++ {
			row[j] = res[i-1][j-1] + res[i-1][j]
		}
		res = append(res, row)
	}
	return res
}

func main() {
	numRows := 5
	fmt.Println(generate(numRows))
}
