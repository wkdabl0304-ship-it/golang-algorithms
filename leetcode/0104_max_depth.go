package main

import "fmt"

type TreeNode0104 struct {
	Val   int
	Left  *TreeNode0104
	Right *TreeNode0104
}

// 标准解法：递归法，缺点是使用系统空间栈
//func maxDepth(root *TreeNode0104) int {
//	if root == nil {
//		return 0
//	}
//	leftDepth := maxDepth(root.Left)
//	rightDepth := maxDepth(root.Right)
//	if leftDepth > rightDepth {
//		return leftDepth + 1
//	}
//	return rightDepth + 1
//}

// 标准解法：层序遍历（迭代法），使用堆空间更安全
func maxDepth(root *TreeNode0104) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode0104{root}
	res := 0
	for len(queue) > 0 {
		levelNum := len(queue)
		for i := 0; i < levelNum; i++ {
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}
			queue = queue[1:]
		}
		res++
	}
	return res
}

func main() {
	a1, a11, a12, a21, a22 := new(TreeNode0104), new(TreeNode0104), new(TreeNode0104), new(TreeNode0104), new(TreeNode0104)
	a1.Left, a1.Right = a11, a12
	a12.Left, a12.Right = a21, a22
	a1.Val, a11.Val, a12.Val, a21.Val, a22.Val = 3, 9, 20, 15, 7
	fmt.Println(maxDepth(a1))
}
