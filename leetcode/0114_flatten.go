package main

import "fmt"

type TreeNode0114 struct {
	Val   int
	Left  *TreeNode0114
	Right *TreeNode0114
}

// 标准解法：逆向后序遍历
func flatten(root *TreeNode0114) {
	previous := (*TreeNode0114)(nil)
	var order func(node *TreeNode0114)
	order = func(node *TreeNode0114) {
		if node == nil {
			return
		}
		order(node.Right)
		order(node.Left)
		node.Left = nil
		node.Right = previous
		previous = node
	}
	order(root)
}

// 思路总结：
// 1.一般层序和中序用迭代，其它一般用递归
// 2.后序遍历过程中左右子树顺序互换，再反转遍历，就能得到前序遍历，反之亦然

func main() {
	a11, a21, a22, a31, a32, a33 := new(TreeNode0114), new(TreeNode0114), new(TreeNode0114), new(TreeNode0114), new(TreeNode0114), new(TreeNode0114)
	a11.Val, a21.Val, a22.Val, a31.Val, a32.Val, a33.Val = 1, 2, 5, 3, 4, 6
	a11.Left, a11.Right = a21, a22
	a21.Left, a21.Right = a31, a32
	a22.Right = a33
	flatten(a11)
	for p := a11; p != nil; p = p.Right {
		fmt.Print(p.Val, " ")
	}
}
