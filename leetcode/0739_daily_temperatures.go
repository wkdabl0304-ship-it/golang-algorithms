package main

import "fmt"

// 标准解法：单调栈
func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	var stack []int
	for i, v := range temperatures {
		for len(stack) > 0 && v > temperatures[stack[len(stack)-1]] {
			current := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[current] = i - current
		}
		stack = append(stack, i)
	}
	return res
}

// 思路总结：原本双重循环是主动去找，改成单调栈后是先放着不管，去遍历后面的

func main() {
	temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println(dailyTemperatures(temperatures))
}
