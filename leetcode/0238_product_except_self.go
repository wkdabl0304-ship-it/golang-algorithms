package main

import "fmt"

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	leftSum, rightSum := 1, 1
	res[0], res[len(res)-1] = 1, 1
	for i := 1; i < len(nums); i++ {
		leftSum *= nums[i-1]
		res[i] = leftSum
	}
	for i := len(nums) - 2; i >= 0; i-- {
		rightSum *= nums[i+1]
		res[i] *= rightSum
	}
	return res
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf(nums))
}
