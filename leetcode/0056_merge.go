package main

import (
	"fmt"
	"slices"
)

func merge(intervals [][]int) [][]int {
	// 1.排序
	slices.SortFunc(intervals, func(a, b []int) int {
		return a[0] - b[0]
	})
	// 2.合并
	res := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		if res[len(res)-1][1] >= intervals[i][0] {
			if intervals[i][1] > res[len(res)-1][1] {
				res[len(res)-1][1] = intervals[i][1]
			}
		} else {
			res = append(res, intervals[i])
		}
	}
	return res
}

func main() {
	//intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}}
	fmt.Println(merge(intervals))
}
