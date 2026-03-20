package main

import "fmt"

type TreeNode0199 struct {
	Val   int
	Left  *TreeNode0199
	Right *TreeNode0199
}

// 个人解法：层序遍历
func rightSideView(root *TreeNode0199) []int {
	if root == nil {
		return nil
	}
	queue := []*TreeNode0199{root}
	var res []int
	for len(queue) > 0 {
		levelNodeNum := len(queue)
		for i := 0; i < levelNodeNum; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			if i+1 == levelNodeNum {
				res = append(res, node.Val)
			}
		}
	}
	return res
}

func main() {
	a11, a21, a22, a31, a32 := new(TreeNode0199), new(TreeNode0199), new(TreeNode0199), new(TreeNode0199), new(TreeNode0199)
	a11.Left, a11.Right = a21, a22
	a21.Right = a31
	a22.Right = a32
	a11.Val, a21.Val, a22.Val, a31.Val, a32.Val = 1, 2, 3, 5, 4
	fmt.Println(rightSideView(a11))
}
