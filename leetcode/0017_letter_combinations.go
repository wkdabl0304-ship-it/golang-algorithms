package main

import "fmt"

// 标准解法：回溯算法
func letterCombinations(digits string) []string {
	var res []string
	var path []byte
	mapping := []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	var backtrack func(index int)
	backtrack = func(index int) {
		if len(path) == len(digits) {
			res = append(res, string(path))
			return
		}
		letters := mapping[digits[index]-'0']
		for i := 0; i < len(letters); i++ {
			path = append(path, letters[i])
			backtrack(index + 1)
			path = path[:len(path)-1]
		}
	}
	backtrack(0)
	return res
}

// 思路总结：
// 1.频繁操作字符串首选类型 []byte
// 2.为了思路清晰，一步一步写

func main() {
	digits := "23"
	fmt.Println(letterCombinations(digits))
}
