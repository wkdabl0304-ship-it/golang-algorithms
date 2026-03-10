package main

import "fmt"

// 我的解法：二次遍历
//func twoSum(nums []int, target int) []int {
//	m := make(map[int]int)
//	for i, v := range nums {
//		m[v] = i
//	}
//	for i, v := range nums {
//		if res, ok := m[target-v]; ok {
//			if res != i {
//				return []int{i, res}
//			}
//		}
//	}
//	return nil
//}

// 最优解法：一次遍历
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		if res, ok := m[target-v]; ok {
			return []int{res, i}
		}
		m[v] = i
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}
