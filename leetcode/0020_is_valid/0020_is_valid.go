package main

import "fmt"

// 1.个人版本(个人觉得更简单易懂，无处理特殊情况，处理时间恒定)
//func isValid(s string) bool {
//	var stack []byte
//	for i := 0; i < len(s); i++ {
//		if len(stack) == 0 {
//			stack = append(stack, s[i])
//		} else {
//			if s[i] == ')' && stack[len(stack)-1] == '(' ||
//				s[i] == '}' && stack[len(stack)-1] == '{' ||
//				s[i] == ']' && stack[len(stack)-1] == '[' {
//				stack = stack[:len(stack)-1]
//			} else {
//				stack = append(stack, s[i])
//			}
//		}
//	}
//	return len(stack) == 0
//}

// 2.个人改良版
//func isValid(s string) bool {
//	n := len(s)
//	if n%2 != 0 {
//		return false
//	}
//	var stack []byte
//	for i := 0; i < len(s); i++ {
//		if s[i] == '[' || s[i] == '{' || s[i] == '(' {
//			stack = append(stack, s[i])
//			continue
//		}
//		if len(stack) == 0 {
//			return false
//		}
//		if (s[i] == '}' && stack[len(stack)-1] != '{') ||
//			(s[i] == ']' && stack[len(stack)-1] != '[') ||
//			(s[i] == ')' && stack[len(stack)-1] != '(') {
//			return false
//		}
//		stack = stack[:len(stack)-1]
//	}
//	return len(stack) == 0
//}

// 3.标准解法(节省时间：1.奇数个字符 2.括号乱序 3.不匹配)
func isValid(s string) bool {
	n := len(s)
	if n%2 != 0 {
		return false
	}
	pairs := map[byte]byte{
		'{': '}',
		'[': ']',
		'(': ')',
	}
	var stack []byte
	for i := 0; i < len(s); i++ {
		if _, ok := pairs[s[i]]; ok {
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 {
				return false
			} else if s[i] != pairs[stack[len(stack)-1]] {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}
	return len(stack) == 0
}

func main() {
	s := "([])"
	res := isValid(s)
	fmt.Println(res)
}
