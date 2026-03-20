package main

import "fmt"

type TreeNode0543 struct {
	Val   int
	Left  *TreeNode0543
	Right *TreeNode0543
}

// 标准解法：递归法，相比于迭代法更简单，迭代法非常复杂且耗费更多空间
func diameterOfBinaryTree(root *TreeNode0543) int {
	maxDiameter := 0
	var depth func(*TreeNode0543) int
	depth = func(node *TreeNode0543) int {
		if node == nil {
			return 0
		}
		leftDepth := depth(node.Left)
		rightDepth := depth(node.Right)
		if leftDepth+rightDepth > maxDiameter {
			maxDiameter = leftDepth + rightDepth
		}
		if leftDepth > rightDepth {
			return leftDepth + 1
		}
		return rightDepth + 1
	}
	depth(root)
	return maxDiameter
}

// 思路总结：
// 1.使用递归并且需要全局变量的时候可以用闭包
// 2.自顶向下使用迭代法（层序遍历），自下而上使用递归法

func main() {
	a11, a21, a22, a31, a32 := new(TreeNode0543), new(TreeNode0543), new(TreeNode0543), new(TreeNode0543), new(TreeNode0543)
	a11.Val, a21.Val, a22.Val, a31.Val, a32.Val = 1, 2, 3, 4, 5
	a11.Left, a11.Right = a21, a22
	a21.Left, a21.Right = a31, a32
	fmt.Println(diameterOfBinaryTree(a11))
}
