package main

import "fmt"

// 标准解法：贪心算法
func partitionLabels(s string) []int {
	n := len(s)
	m := make([]int, 26)
	for i := 0; i < n; i++ {
		m[s[i]-'a'] = i
	}
	var res []int
	start, end := 0, 0
	for i := 0; i < n; i++ {
		if m[s[i]-'a'] > end {
			end = m[s[i]-'a']
		}
		if i == end {
			res = append(res, end-start+1)
			start = i + 1
		}
	}
	return res
}

// 思路总结：
// 1.如果题目提到字符串的字符全是小写字母的时候，为了优化可以用 [26]int，包括大写字母的话用 [58]int（大写字母ASCII码更小）
// 2.如果发现一些数据不足的时候，要考虑是否先遍历一遍数组再说

func main() {
	s := "ababcbacadefegdehijhklij"
	fmt.Println(partitionLabels(s))
}
