package main

import "fmt"

// MinStack 标准解法：双栈
type MinStack struct {
	stack  []int
	minNum []int
}

func Constructor0155() MinStack {
	return MinStack{}
}

func (s *MinStack) Push(val int) {
	s.stack = append(s.stack, val)
	if len(s.minNum) == 0 {
		s.minNum = append(s.minNum, val)
		return
	}
	current := s.minNum[len(s.minNum)-1]
	if val < current {
		s.minNum = append(s.minNum, val)
	} else {
		s.minNum = append(s.minNum, current)
	}
}

func (s *MinStack) Pop() {
	s.stack = s.stack[:len(s.stack)-1]
	s.minNum = s.minNum[:len(s.minNum)-1]
}

func (s *MinStack) Top() int {
	return s.stack[len(s.stack)-1]
}

func (s *MinStack) GetMin() int {
	return s.minNum[len(s.minNum)-1]
}

// 思路总结：
// 1.要最小值，首先想到用一个变量记录最小值
// 2.但是如果最小值出栈了，并不知道上一个最小值是多少，所以得用一个数据结构能存储所有最小值
// 3.但是并不是所有的值都需要存储，只要存储当时栈元素对应的最小值就好了
// 4.所以引入第二个栈（其实用一个栈但是双int数据结构也可以）
func main() {
	obj := Constructor0155()
	obj.Push(-3)
	obj.Push(-2)
	fmt.Println(obj.GetMin())
	fmt.Println(obj.Top())
	obj.Pop()
	fmt.Println(obj.Top())
	// -3 -2 -3
}
