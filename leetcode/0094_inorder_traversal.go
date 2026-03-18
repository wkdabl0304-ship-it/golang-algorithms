package main

import "fmt"

type TreeNode0094 struct {
	Val   int
	Left  *TreeNode0094
	Right *TreeNode0094
}

// 个人解法：递归法，写法简单，易于理解
//func inorderTraversal(root *TreeNode0094) []int {
//	var res []int
//	var dfs func(*TreeNode0094)
//	dfs = func(node *TreeNode0094) {
//		if node == nil {
//			return
//		}
//		dfs(node.Left)
//		res = append(res, node.Val)
//		dfs(node.Right)
//	}
//	dfs(root)
//	return res
//}

// 标准解法：迭代法，可控、操作自由并且不会受限于系统栈空间
func inorderTraversal(root *TreeNode0094) []int {
	var res []int
	current := root
	var stack []*TreeNode0094
	for current != nil || len(stack) > 0 {
		// 1.一路往左走
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}
		// 2.记录和弹栈
		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, current.Val)
		// 3.往右走
		current = current.Right
	}
	return res
}

// 思路总结：每个循环内都有向左下走的小循环，大循环是负责回退

func main() {
	a1, a2, a3 := new(TreeNode0094), new(TreeNode0094), new(TreeNode0094)
	a1.Val, a2.Val, a3.Val = 1, 2, 3
	a1.Right = a2
	a2.Left = a3
	fmt.Println(inorderTraversal(a1))
}
