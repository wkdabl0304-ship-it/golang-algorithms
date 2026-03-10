package main

import "fmt"

// 个人解法：时间复杂度O(n2)
//func longestConsecutive(nums []int) int {
//	m := make(map[int]bool)
////	for _, v := range nums {
////		m[v] = true
////	}
////	res := 0
//	for _, v := range nums {
//		cnt := 1
//		for m[v+cnt] {
//			cnt++
//		}
//		if cnt > res {
//			res = cnt
//		}
//	}
//	return res
//}

// 标准解法：只遍历是首位数的，逻辑上时间复杂度为O(n)
func longestConsecutive(nums []int) int {
	m := make(map[int]bool)
	for _, v := range nums {
		m[v] = true
	}
	res := 0
	for k := range m {
		if !m[k-1] {
			cnt := 1
			for m[k+cnt] {
				cnt++
			}
			if cnt > res {
				res = cnt
			}
		}
	}
	return res
}

func main() {
	nums := []int{100, 4, 200, 1, 3, 2}
	fmt.Println(longestConsecutive(nums))
}
