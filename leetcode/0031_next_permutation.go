package main

import "fmt"

// 标准解法：双指针
func nextPermutation(nums []int) {
	n := len(nums)
	i, j := n-2, n-1
	for ; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			break
		}
	}
	if i >= 0 {
		for ; j > i; j-- {
			if nums[j] > nums[i] {
				break
			}
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	reverse(nums[i+1:])
}

func reverse(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

// 思路总结：排列的下一个排列指的是将尾部递减序前的数字与递减序里的第一个比其大的数字交换，再反转序列

func main() {
	nums := []int{1, 2, 7, 4, 3, 1}
	nextPermutation(nums)
	fmt.Println(nums)
}
