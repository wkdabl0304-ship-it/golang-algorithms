package main

import "fmt"

type TreeNode0101 struct {
	Val   int
	Left  *TreeNode0101
	Right *TreeNode0101
}

// 个人解法：为了防止混淆，每一种情况都罗列出来
//func isSymmetric(root *TreeNode0101) bool {
//	queue := []*TreeNode0101{root.Left, root.Right}
//	for len(queue) > 0 {
//		node1, node2 := queue[0], queue[1]
//		// 1.一空一非空
//		if (node1 == nil && node2 != nil) || (node1 != nil && node2 == nil) {
//			return false
//		}
//		// 2.都非空
//		if node1 != nil && node2 != nil {
//			if node1.Val != node2.Val {
//				// 值不相等
//				return false
//			} else {
//				// 值相等
//				queue = queue[2:]                                                       // 出队
//				queue = append(queue, node1.Left, node2.Right, node1.Right, node2.Left) // 入队
//			}
//		}
//		// 3.都空
//		if node1 == nil && node2 == nil {
//			queue = queue[2:]
//		}
//	}
//	return true
//}

// 标准解法：通过有技巧地分类，将代码简化
func isSymmetric(root *TreeNode0101) bool {
	queue := []*TreeNode0101{root.Left, root.Right}
	for len(queue) > 0 {
		node1, node2 := queue[0], queue[1]
		queue = queue[2:]
		if node1 == nil && node2 == nil {
			continue
		}
		if node1 == nil || node2 == nil || node1.Val != node2.Val {
			return false
		}
		queue = append(queue, node1.Left, node2.Right, node1.Right, node2.Left)
	}
	return true
}

func main() {
	a11, a21, a22, a31, a32, a33, a34 := new(TreeNode0101), new(TreeNode0101), new(TreeNode0101), new(TreeNode0101), new(TreeNode0101), new(TreeNode0101), new(TreeNode0101)
	a11.Val, a21.Val, a22.Val, a31.Val, a32.Val, a33.Val, a34.Val = 1, 2, 2, 3, 4, 4, 3
	a11.Left, a11.Right = a21, a22
	a21.Left, a21.Right = a31, a32
	a22.Left, a22.Right = a33, a34
	fmt.Println(isSymmetric(a11))
}
