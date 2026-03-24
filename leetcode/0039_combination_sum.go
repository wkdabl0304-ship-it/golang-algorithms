package main

import (
	"fmt"
	"slices"
)

// 标准解法：回溯算法
func combinationSum(candidates []int, target int) [][]int {
	slices.Sort(candidates)
	var path []int
	currentSum := 0
	var res [][]int
	var backtrack func(index int)
	backtrack = func(index int) {
		if currentSum == target {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}
		for i := index; i < len(candidates); i++ {
			if currentSum+candidates[i] > target {
				break
			}
			currentSum += candidates[i]
			path = append(path, candidates[i])
			backtrack(i)
			currentSum -= candidates[i]
			path = path[:len(path)-1]
		}
	}
	backtrack(0)
	return res
}

// 思路总结：
// 1.固定个数的情况需要条件判断，否则不需要
// 2.递归时，组合需要传参，排列不需要传参

func main() {
	candidates := []int{2, 3, 6, 7}
	target := 7
	fmt.Println(combinationSum(candidates, target))
}
