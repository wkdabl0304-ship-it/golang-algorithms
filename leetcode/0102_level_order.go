package main

import "fmt"

type TreeNode0102 struct {
	Val   int
	Left  *TreeNode0102
	Right *TreeNode0102
}

// 个人解法 & 标准解法：队列 + 数量统计分层
func levelOrder(root *TreeNode0102) [][]int {
	if root == nil {
		return nil
	}
	var res [][]int
	queue := []*TreeNode0102{root}
	for len(queue) > 0 {
		levelNodeNum := len(queue)
		levelArr := make([]int, 0, levelNodeNum)
		for i := 0; i < levelNodeNum; i++ {
			node := queue[0]
			queue = queue[1:]
			levelArr = append(levelArr, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, levelArr)
	}
	return res
}

func main() {
	a11, a21, a22, a31, a32 := new(TreeNode0102), new(TreeNode0102), new(TreeNode0102), new(TreeNode0102), new(TreeNode0102)
	a11.Left, a11.Right = a21, a22
	a22.Left, a22.Right = a31, a32
	a11.Val, a21.Val, a22.Val, a31.Val, a32.Val = 3, 9, 20, 15, 7
	fmt.Println(levelOrder(a11))
}
