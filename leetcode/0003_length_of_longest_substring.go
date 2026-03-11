package main

import "fmt"

// 个人解法：暴力回退导致时间复杂度为O(n2)，其次是频繁创建 map 导致内存损耗严重
//func lengthOfLongestSubstring(s string) int {
//	res := 0
//	if len(s) == 0 {
//		return 0
//	}
//	left, right := 0, 0
//	for left < len(s) {
//		m := make(map[byte]int)
//		for right < len(s) {
//			if v, ok := m[s[right]]; !ok {
//				m[s[right]] = right
//				right++
//			} else {
//				left = v + 1
//				right = left
//				if len(m) > res {
//					res = len(m)
//				}
//				clear(m)
//				break
//			}
//		}
//		if right == len(s) {
//			if len(m) > res {
//				res = len(m)
//			}
//			break
//		}
//	}
//	return res
//}

// 标准解法：真正的滑动窗口，不回退，通过维护 map 来复用 map
func lengthOfLongestSubstring(s string) int {
	res := 0
	slow := 0
	m := make(map[byte]int)
	for fast := 0; fast < len(s); fast++ {
		if i, ok := m[s[fast]]; ok && i >= slow {
			slow = i + 1
		}
		m[s[fast]] = fast
		if fast-slow+1 > res {
			res = fast - slow + 1
		}
	}
	return res
}

// 经验总结：
// 1.一般快指针不会回退，以快指针做循环，让慢指针去追快指针
// 2.易错点：出现重复字符的时候要保证在当前子串内
// 3.写算法的时候去具体想每一步的时候，就很容易写出来

func main() {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s))
}
