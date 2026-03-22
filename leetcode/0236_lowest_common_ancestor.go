package main

import "fmt"

type TreeNode0236 struct {
	Val   int
	Left  *TreeNode0236
	Right *TreeNode0236
}

// 标准解法：后序遍历
func lowestCommonAncestor(root, p, q *TreeNode0236) *TreeNode0236 {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	return right
}

// 思路总结：面对需要自底向上的题目，一般使用后序遍历

func main() {
	a11, a21, a22, a31, a32, a33, a34, a41, a42 := new(TreeNode0236), new(TreeNode0236), new(TreeNode0236), new(TreeNode0236), new(TreeNode0236), new(TreeNode0236), new(TreeNode0236), new(TreeNode0236), new(TreeNode0236)
	a11.Left, a11.Right = a21, a22
	a21.Left, a21.Right = a31, a32
	a22.Left, a22.Right = a33, a34
	a32.Left, a32.Right = a41, a42
	a11.Val, a21.Val, a22.Val, a31.Val, a32.Val, a33.Val, a34.Val, a41.Val, a42.Val = 3, 5, 1, 6, 2, 0, 8, 7, 4
	fmt.Println(lowestCommonAncestor(a11, a21, a22).Val)
}
