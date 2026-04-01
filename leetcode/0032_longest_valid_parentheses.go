package main

import "fmt"

// 标准解法：栈
func longestValidParentheses(s string) int {
	stack := []int{0: -1}
	res := 0
	n := len(s)
	for i := 0; i < n; i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else if s[i] == ')' {
			stack = stack[:len(stack)-1]
			if len(stack) > 0 {
				index := stack[len(stack)-1]
				if i-index > res {
					res = i - index
				}
			} else {
				stack = append(stack, i)
			}
		}
	}
	return res
}

// 思路总结：遇到括号匹配问题就用栈，八九不离十，只是如何变换的问题

func main() {
	s := ")()())"
	fmt.Println(longestValidParentheses(s))
}
