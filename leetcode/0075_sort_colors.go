package main

import "fmt"

// 标准解法：三指针
func sortColors(nums []int) {
	index, left, right := 0, 0, len(nums)-1
	for index <= right {
		if nums[index] == 0 {
			nums[left], nums[index] = nums[index], nums[left]
			left++
			index++
		} else if nums[index] == 2 {
			nums[right], nums[index] = nums[index], nums[right]
			right--
		} else {
			index++
		}
	}
}

// 思路总结：
// 左指针和右指针功能并不完全一致，因为 left 指向的一定是1，所以只需交换一次，但 right 未知

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	fmt.Println(nums)
}
