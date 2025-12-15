package main

import "fmt"

// 心得体会
// 1.把零移到后面太难了就可以反过来思考，把非零的元素移到前面

func moveZeroes(nums []int) {

	for slow, fast := 0, 0; fast < len(nums); fast++ {
		if nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
	}

}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	//nums := []int{0, 0}
	//nums := []int{1, 0}
	//nums := []int{1, 0, 1}
	moveZeroes(nums)
	fmt.Println(nums)
}
