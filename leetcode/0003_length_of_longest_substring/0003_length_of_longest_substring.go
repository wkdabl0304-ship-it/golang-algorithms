package main

import "fmt"

// 心得体会
// 1.维护一个 map 的增删是比较耗时的，可以改成各种条件的判断
// 2.滑动窗口特征：根据条件移动左指针，循环移动右指针

func lengthOfLongestSubstring(s string) int {
	//1.自研删 map 法
	//maxLen := 0
	//m := make(map[byte]int)
	//left, right := 0, 0
	//for right < len(s) {
	//	_, ok := m[s[right]]
	//	if !ok {
	//		m[s[right]] = right
	//		right++
	//		continue
	//	}
	//	length := right - left
	//	if length > maxLen {
	//		maxLen = length
	//	}
	//	same := s[right]
	//	i2 := m[same]
	//	for k := left; k <= i2; k++ {
	//		delete(m, s[k])
	//	}
	//	left = i2 + 1
	//}
	//if right-left > maxLen {
	//	maxLen = right - left
	//}
	//return maxLen

	//2.个人写法
	//maxLen := 0
	//m := make(map[byte]int)
	//left, right := 0, 0
	//for ; right < len(s); right++ {
	//	pos, ok := m[s[right]]
	//	if !ok {
	//		m[s[right]] = right
	//	} else {
	//		if pos >= left {
	//			length := right - left
	//			if length > maxLen {
	//				maxLen = length
	//			}
	//			left = pos + 1
	//		}
	//	}
	//	m[s[right]] = right
	//}
	//if right-left > maxLen {
	//	maxLen = right - left
	//}
	//return maxLen

	//3.标准解法
	m := make(map[byte]int)
	maxLen := 0
	for left, right := 0, 0; right < len(s); right++ {
		if pos, ok := m[s[right]]; ok && pos >= left {
			left = pos + 1
		}
		m[s[right]] = right
		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}
	return maxLen
}

func main() {
	s := "abcabcbb"
	//s := "pwwkew"
	res := lengthOfLongestSubstring(s)
	fmt.Println(res)
}
