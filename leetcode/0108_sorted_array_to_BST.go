package main

type TreeNode0108 struct {
	Val   int
	Left  *TreeNode0108
	Right *TreeNode0108
}

// 个人解法：闭包递归法，标准解法类似，只是没用闭包
func sortedArrayToBST(nums []int) *TreeNode0108 {
	if len(nums) == 0 {
		return nil
	}
	var bst func(left, right int) *TreeNode0108
	bst = func(left, right int) *TreeNode0108 {
		if left > right {
			return nil
		}
		middle := (left + right) / 2
		newNode := new(TreeNode0108)
		newNode.Val = nums[middle]
		newNode.Left = bst(left, middle-1)
		newNode.Right = bst(middle+1, right)
		return newNode
	}
	return bst(0, len(nums)-1)
}

func main() {
	nums := []int{-10, -3, 0, 5, 9}
	sortedArrayToBST(nums)
}
