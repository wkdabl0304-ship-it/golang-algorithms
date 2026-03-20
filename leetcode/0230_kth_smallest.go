package main

import "fmt"

type TreeNode0230 struct {
	Val   int
	Left  *TreeNode0230
	Right *TreeNode0230
}

// 个人做法：递归法
//func kthSmallest(root *TreeNode0230, k int) int {
//	res, cnt := 0, 0
//	var inorder func(node *TreeNode0230)
//	inorder = func(node *TreeNode0230) {
//		if node == nil {
//			return
//		}
//		inorder(node.Left)
//		cnt++
//		if cnt == k {
//			res = node.Val
//		}
//		inorder(node.Right)
//	}
//	inorder(root)
//	return res
//}

// 标准做法：迭代法
func kthSmallest(root *TreeNode0230, k int) int {
	cnt := 0
	current := root
	var stack []*TreeNode0230
	for current != nil || len(stack) > 0 {
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}
		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		cnt++
		if cnt == k {
			return current.Val
		}
		current = current.Right
	}
	return -1
}

// 思路总结：只要是用到中序遍历的都可以用迭代法

func main() {
	a11, a21, a22, a31 := new(TreeNode0230), new(TreeNode0230), new(TreeNode0230), new(TreeNode0230)
	a11.Left, a11.Right = a21, a22
	a21.Right = a31
	a11.Val, a21.Val, a22.Val, a31.Val = 3, 1, 4, 2
	fmt.Println(kthSmallest(a11, 1))
}
