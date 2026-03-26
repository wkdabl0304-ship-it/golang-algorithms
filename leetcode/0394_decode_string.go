package main

import (
	"fmt"
	"strings"
)

// 标准解法：双栈法
func decodeString(s string) string {
	currentStr, currentNum := "", 0
	var numStack []int
	var strStack []string
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			currentNum = currentNum*10 + int(s[i]-'0')
		} else if s[i] == '[' {
			numStack = append(numStack, currentNum)
			strStack = append(strStack, currentStr)
			currentNum = 0
			currentStr = ""
		} else if s[i] == ']' {
			num := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]
			str := strStack[len(strStack)-1]
			strStack = strStack[:len(strStack)-1]
			currentStr = str + strings.Repeat(currentStr, num)
		} else {
			currentStr += string(s[i])
		}
	}
	return currentStr
}

// 思路总结：
// 1.难点是会出现四种情况并且没有先后顺序，需要并列讨论
// 2.双栈记录入栈前的数字和字符串，其实相当于用了递归

func main() {
	s := "2[abc]3[cd]ef"
	fmt.Println(decodeString(s))
}
