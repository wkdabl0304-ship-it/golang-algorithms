package main

import "fmt"

type TreeNode0437 struct {
	Val   int
	Left  *TreeNode0437
	Right *TreeNode0437
}

// 个人解法：双重递归，但是时间复杂度开销为 O(n2)
//func pathSum(root *TreeNode0437, targetSum int) int {
//	if root == nil {
//		return 0
//	}
//	return count(root, targetSum) + pathSum(root.Left, targetSum) + pathSum(root.Right, targetSum)
//}
//
//func count(root *TreeNode0437, targetSum int) int {
//	res := 0
//	var calculate func(node *TreeNode0437, targetSum int)
//	calculate = func(node *TreeNode0437, targetSum int) {
//		if node == nil {
//			return
//		}
//		if node.Val == targetSum {
//			res++
//		}
//		calculate(node.Left, targetSum-node.Val)
//		calculate(node.Right, targetSum-node.Val)
//	}
//	calculate(root, targetSum)
//	return res
//}

// 标准解法：单递归 + 前缀和
func pathSum(root *TreeNode0437, targetSum int) int {
	res := 0
	m := make(map[int]int)
	m[0] = 1
	var dfs func(node *TreeNode0437, currentSum int)
	dfs = func(node *TreeNode0437, currentSum int) {
		if node == nil {
			return
		}
		currentSum += node.Val
		if num, ok := m[currentSum-targetSum]; ok {
			res += num
		}
		m[currentSum]++
		dfs(node.Left, currentSum)
		dfs(node.Right, currentSum)
		m[currentSum]--
	}
	dfs(root, 0)
	return res
}

// 思路总结：一般遇到连续和的问题，可以使用前缀和

func main() {
	a11, a21, a22, a31, a32, a33, a41, a42, a43 := new(TreeNode0437), new(TreeNode0437), new(TreeNode0437), new(TreeNode0437), new(TreeNode0437), new(TreeNode0437), new(TreeNode0437), new(TreeNode0437), new(TreeNode0437)
	a11.Left, a11.Right = a21, a22
	a21.Left, a21.Right = a31, a32
	a22.Right = a33
	a31.Left, a31.Right = a41, a42
	a32.Right = a43
	a11.Val, a21.Val, a22.Val, a31.Val, a32.Val, a33.Val, a41.Val, a42.Val, a43.Val = 10, 5, -3, 3, 2, 11, 3, -2, 1
	fmt.Println(pathSum(a11, 8))
}
