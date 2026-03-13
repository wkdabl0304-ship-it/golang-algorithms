package main

import (
	"fmt"
	"math"
)

// 标准解法：滑动窗口
func minWindow(s string, t string) string {
	nums, need := make(map[byte]int), make(map[byte]int)
	left, right := 0, 0
	start, minLen := 0, math.MaxInt32
	match := 0
	// 1.初始化
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	for right < len(s) {
		// 2.右指针扩张
		if _, ok := need[s[right]]; ok {
			nums[s[right]]++
			if nums[s[right]] == need[s[right]] {
				match++
			}
		}
		right++
		// 3.左指针收缩
		for match == len(need) {
			if _, ok := need[s[left]]; ok {
				if right-left < minLen {
					minLen = right - left
					start = left
				}
				if nums[s[left]] == need[s[left]] {
					match--
				}
				nums[s[left]]--
			}
			left++
		}
	}
	if minLen == math.MaxInt32 {
		return ""
	}
	return s[start : start+minLen]
}

// 解题思路：
// 1.通过 map + match种类数（意味着用 map 更好，而不是数组）来实现是否包含子串的判断
// 2.通过 滑动窗口 快慢指针来避免嵌套循环

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"
	fmt.Println(minWindow(s, t))
}
