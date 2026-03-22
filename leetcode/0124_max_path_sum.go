package main

import (
	"fmt"
	"math"
)

type TreeNode0124 struct {
	Val   int
	Left  *TreeNode0124
	Right *TreeNode0124
}

// 标准解法：后序遍历 + 动态规划，和 LC0053 唯一的区别是它是两头一起进行的
func maxPathSum(root *TreeNode0124) int {
	maxSum := math.MinInt64
	var maxPath func(node *TreeNode0124) int
	maxPath = func(node *TreeNode0124) int {
		if node == nil {
			return 0
		}
		left := max(0, maxPath(node.Left))
		right := max(0, maxPath(node.Right))
		if sum := left + right + node.Val; sum > maxSum {
			maxSum = sum
		}
		return node.Val + max(left, right)
	}
	maxPath(root)
	return maxSum
}

func main() {
	a11, a21, a22, a31, a32 := new(TreeNode0124), new(TreeNode0124), new(TreeNode0124), new(TreeNode0124), new(TreeNode0124)
	a11.Left, a11.Right = a21, a22
	a22.Left, a22.Right = a31, a32
	a11.Val, a21.Val, a22.Val, a31.Val, a32.Val = -10, 9, 20, 15, 7
	fmt.Println(maxPathSum(a11))
}
