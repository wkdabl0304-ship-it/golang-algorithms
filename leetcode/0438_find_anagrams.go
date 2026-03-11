package main

import "fmt"

// 个人解法（标准解法）：和《异位词分组》题目判断异位词的方式一致，滑动固定大小的窗口
func findAnagrams(s string, p string) []int {
	sLen, pLen := len(s), len(p)
	if sLen < pLen {
		return nil
	}
	var res []int
	var standard, char [26]int
	for i := 0; i < pLen; i++ {
		standard[p[i]-'a']++
		char[s[i]-'a']++
	}
	slow, fast := 0, pLen-1
	for fast < sLen {
		if char == standard {
			res = append(res, slow)
		}
		fast++
		if fast == sLen {
			break
		}
		char[s[fast]-'a']++
		char[s[slow]-'a']--
		slow++
	}
	return res
}

func main() {
	s := "cbaebabacd"
	p := "abc"
	fmt.Println(findAnagrams(s, p))
}
