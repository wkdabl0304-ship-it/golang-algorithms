package main

import "fmt"

// 心得体会
// 1.当要反复查询一个数组的时候可以将数组转 map

func twoSum(nums []int, target int) []int {

	//1.暴力遍历法
	//for index, value := range nums {
	//	for i := index + 1; i < len(nums); i++ {
	//		if value + nums[i] == target {
	//			return []int{index,i}
	//		}
	//	}
	//}

	//2.map 改良（两次遍历）
	//getNum := make(map[int]int, len(nums))
	//getIndex := make(map[int]int, len(nums))
	//for i, v := range nums {
	//	getNum[i] = v
	//	getIndex[v] = i
	//}
	//for i, v := range nums {
	//	targetNum := target - v
	//	targetIndex, ok := getIndex[targetNum]
	//	if ok && i != targetIndex {
	//		return []int{i, targetIndex}
	//	}
	//}

	//3.完整标准解法
	//pn := make(map[int]int, len(nums))
	//for i, v := range nums {
	//	targetNum := target - v
	//	targetIndex, ok := pn[targetNum]
	//	if ok {
	//		return []int{targetIndex, i}
	//	} else {
	//		pn[v] = i
	//	}
	//}

	//4.最简标准解法
	pn := make(map[int]int, len(nums))
	for i, v := range nums {
		if ti, ok := pn[target-v]; ok {
			return []int{ti, i}
		}
		pn[v] = i
	}

	return nil
}

func main() {

	nums := []int{3, 2, 4}

	target := 6

	res := twoSum(nums, target)

	fmt.Println(res)

}
