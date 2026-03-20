package main

import (
	"fmt"
	"math"
)

type TreeNode0098 struct {
	Val   int
	Left  *TreeNode0098
	Right *TreeNode0098
}

// 标准解法：中序遍历
func isValidBST(root *TreeNode0098) bool {
	previous := math.MinInt64
	var inorder func(node *TreeNode0098) bool
	inorder = func(node *TreeNode0098) bool {
		if node == nil {
			return true
		}
		if !inorder(node.Left) {
			return false
		}
		if node.Val <= previous {
			return false
		}
		previous = node.Val
		return inorder(node.Right)
	}
	return inorder(root)
}

// 思路总结：二叉搜索树在中序遍历下满足严格单调递增

func main() {
	a11, a21, a22, a31, a32 := new(TreeNode0098), new(TreeNode0098), new(TreeNode0098), new(TreeNode0098), new(TreeNode0098)
	a11.Left, a11.Right = a21, a22
	a22.Left, a22.Right = a31, a32
	a11.Val, a21.Val, a22.Val, a31.Val, a32.Val = 5, 1, 4, 3, 6
	fmt.Println(isValidBST(a11))
}
