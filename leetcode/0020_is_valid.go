package main

import "fmt"

// 个人解法 & 标准解法：栈
func isValid(s string) bool {
	var stack []int32
	for _, v := range s {
		if v == '{' || v == '(' || v == '[' {
			stack = append(stack, v)
			continue
		}
		if len(stack) == 0 {
			return false
		}
		if (v == '}' && stack[len(stack)-1] == '{') ||
			(v == ']' && stack[len(stack)-1] == '[') ||
			(v == ')' && stack[len(stack)-1] == '(') {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}

func main() {
	s := "()[]{}"
	fmt.Println(isValid(s))
}
