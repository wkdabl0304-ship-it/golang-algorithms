package main

import "fmt"

// 心得体会
// 1.如果只需要比较两个字符串(或数组)含有的元素个数是否相同，可以用差值计算
// 2.这题的精髓就是维护一个 count 来区分两个字符串是否相同(取代字符串的遍历)

// 个人分析：
// 1.判断是否是异位词的条件：子串中字母出现的次数和目标字符串一致 -> 用 map 来统计出现次数
// 2.子串长度恒等于目标字符串长度 -> 固定大小的滑动窗口
//func findAnagrams(s string, p string) []int {
//	var res []int
//
//	// 0.异常情况
//	if len(s) < len(p) {
//		return res
//	}
//
//	// 1.目标串的字符统计 && 字串统计初始化
//	m1 := make(map[byte]int, 26) // 目标字符串
//	m2 := make(map[byte]int, 26) // 子串
//	for i := 0; i < len(p); i++ {
//		if _, ok := m1[p[i]]; !ok {
//			m1[p[i]] = 0
//		}
//		m1[p[i]]++
//		if _, ok := m2[s[i]]; !ok {
//			m2[s[i]] = 0
//		}
//		m2[s[i]]++
//	}
//
//	// 2.核心过程
//	left, right := 0, len(p)-1
//	for right < len(s) {
//		// 判断相等
//		equal := true
//		for k, v := range m1 {
//			if m2[k] != v {
//				equal = false
//				break
//			}
//		}
//		if equal {
//			res = append(res, left)
//		}
//
//		// 收尾工作
//		m2[s[left]]--
//		left++
//		right++
//		if right == len(s) {
//			break
//		}
//		if _, ok := m2[s[right]]; ok {
//			m2[s[right]]++
//		} else {
//			m2[s[right]] = 1
//		}
//	}
//
//	return res
//}

// 标准解法
//func findAnagrams(s string, p string) []int {
//	res := []int{}
//	if len(s) < len(p) {
//		return res
//	}
//
//	need := make(map[byte]int)
//	for i := 0; i < len(p); i++ {
//		need[p[i]]++
//	}
//
//	left, right := 0, 0
//	needCount := len(p)
//
//	for right < len(s) {
//		if need[s[right]] > 0 {
//			needCount--
//		}
//		need[s[right]]--
//		right++
//
//		if right-left == len(p) {
//			if needCount == 0 {
//				res = append(res, left)
//			}
//			if need[s[left]] >= 0 {
//				needCount++
//			}
//			need[s[left]]++
//			left++
//		}
//	}
//
//	return res
//}

// 个人版本
// 个人觉得更好理解，standard 记录子串根据目标字符串的相对差异(多了还是少了)，need 记录需要的总额
func findAnagrams(s string, p string) []int {
	var res []int
	var standard [26]int
	need := len(p)
	for i := 0; i < len(p); i++ {
		standard[p[i]-'a']--
	}
	for left, right := 0, 0; right < len(s); {
		if standard[s[right]-'a'] < 0 {
			need--
		}
		standard[s[right]-'a']++
		right++

		if right-left == len(p) {
			if need == 0 {
				res = append(res, left)
			}
			if standard[s[left]-'a'] <= 0 {
				need++
			}
			standard[s[left]-'a']--
			left++
		}
	}
	return res
}

func main() {
	s, p := "cbaebabacd", "abc"
	res := findAnagrams(s, p)
	fmt.Println(res)
}
