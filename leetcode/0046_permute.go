package main

import "fmt"

// 标准解法：回溯算法
func permute(nums []int) [][]int {
	used := make([]bool, len(nums))
	path := make([]int, 0, len(nums))
	var res [][]int
	var backtrack func()
	backtrack = func() {
		if len(path) == len(nums) {
			temp := make([]int, len(nums))
			copy(temp, path)
			res = append(res, temp)
		}
		for i, v := range nums {
			if used[i] {
				continue
			}
			path = append(path, v)
			used[i] = true
			backtrack()
			path = path[:len(path)-1]
			used[i] = false
		}
	}
	backtrack()
	return res
}

// 思路总结：
// 1.因为要不断更新当前路径，所以一般采用闭包
// 2.因为要区分是否走过，所以需要标记
// 3.因为要遍历所有情况，所以要回退
// 4.回退前要记得清除标记

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(permute(nums))
}
