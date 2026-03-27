package main

import "fmt"

// 标准解法：回溯算法
func partition0131(s string) [][]string {
	var path []string
	var res [][]string
	var backtrack func(start int)
	backtrack = func(start int) {
		if start == len(s) {
			temp := make([]string, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}
		for i := start; i < len(s); i++ {
			subString := s[start : i+1]
			if isPalindrome0131(subString) {
				path = append(path, subString)
				backtrack(i + 1)
				path = path[:len(path)-1]
			}
		}
	}
	backtrack(0)
	return res
}

func isPalindrome0131(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func main() {
	s := "aab"
	fmt.Println(partition0131(s))
}
