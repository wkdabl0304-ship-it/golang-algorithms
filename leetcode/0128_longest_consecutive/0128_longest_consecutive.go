package main

import "fmt"

// 心得体会
// 1.其实循环+记录状态可以代替递归
// 2.map 除了加速查找的作用，还有去重的作用
// 3.golang 实现集合的方法可用 map[any]bool

func longestConsecutive(nums []int) int {

	if len(nums) == 0 {
		return 0
	}
	m := make(map[int]bool)
	for _, v := range nums {
		m[v] = true
	}
	maxLen := 0
	for k := range m {
		if !m[k-1] {
			length := 1
			cur := k
			for m[cur+1] {
				length++
				cur++
			}
			if length > maxLen {
				maxLen = length
			}
		}
	}
	return maxLen

}

func main() {
	nums := []int{100, 4, 200, 1, 3, 2}

	res := longestConsecutive(nums)

	fmt.Println(res)
}
