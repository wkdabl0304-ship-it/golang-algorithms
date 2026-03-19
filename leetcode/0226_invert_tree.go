package main

import "fmt"

type TreeNode0226 struct {
	Val   int
	Left  *TreeNode0226
	Right *TreeNode0226
}

// 标准解法：层序遍历（迭代法）
func invertTree(root *TreeNode0226) *TreeNode0226 {
	if root == nil {
		return nil
	}
	queue := []*TreeNode0226{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		node.Left, node.Right = node.Right, node.Left
	}
	return root
}

func main() {
	a11, a21, a22, a31, a32, a33, a34 := new(TreeNode0226), new(TreeNode0226), new(TreeNode0226), new(TreeNode0226), new(TreeNode0226), new(TreeNode0226), new(TreeNode0226)
	a11.Left, a11.Right = a21, a22
	a21.Left, a21.Right = a31, a32
	a22.Left, a22.Right = a33, a34
	a11.Val, a21.Val, a22.Val, a31.Val, a32.Val, a33.Val, a34.Val = 4, 2, 7, 1, 3, 6, 9
	fmt.Println(invertTree(a11))
}
