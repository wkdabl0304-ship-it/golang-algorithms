package main

import "fmt"

// 标准解法：回溯算法
func generateParenthesis(n int) []string {
	var res []string
	var path []byte
	leftCount, rightCount := 0, 0
	var backtrack func()
	backtrack = func() {
		if leftCount == rightCount && leftCount == n {
			res = append(res, string(path))
		}
		// 左括号情况
		if leftCount < n {
			leftCount++
			path = append(path, '(')
			backtrack()
			path = path[:len(path)-1]
			leftCount--
		}
		// 右括号情况
		if rightCount < leftCount && rightCount < n {
			rightCount++
			path = append(path, ')')
			backtrack()
			path = path[:len(path)-1]
			rightCount--
		}
	}
	backtrack()
	return res
}

func main() {
	n := 3
	fmt.Println(generateParenthesis(n))
}
