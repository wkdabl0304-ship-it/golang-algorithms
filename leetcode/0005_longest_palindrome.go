package main

import "fmt"

// 标准解法：双指针中心扩散
func longestPalindrome(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}
	res := ""
	for i := 0; i < n; i++ {
		s1 := expand(s, i, i)
		s2 := expand(s, i, i+1)
		if len(s1) > len(res) {
			res = s1
		}
		if len(s2) > len(res) {
			res = s2
		}
	}
	return res
}

func expand(s string, left, right int) string {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return s[left+1 : right]
}

// 思路总结：回文字符串依然用双指针，只是看如何变换

func main() {
	s := "babad"
	fmt.Println(longestPalindrome(s))
}
