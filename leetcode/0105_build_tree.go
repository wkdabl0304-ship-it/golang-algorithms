package main

type TreeNode0105 struct {
	Val   int
	Left  *TreeNode0105
	Right *TreeNode0105
}

// 标准解法：前序遍历 + 中序遍历
func buildTree(preorder []int, inorder []int) *TreeNode0105 {
	m := make(map[int]int, len(preorder))
	for i := 0; i < len(preorder); i++ {
		m[inorder[i]] = i
	}
	var build func(preLeft, preRight, inLeft, inRight int) *TreeNode0105
	build = func(preLeft, preRight, inLeft, inRight int) *TreeNode0105 {
		if preLeft > preRight {
			return nil
		}
		root := new(TreeNode0105)
		root.Val = preorder[preLeft]
		rootIndex := m[preorder[preLeft]]
		leftSize := rootIndex - inLeft
		root.Left = build(preLeft+1, preLeft+leftSize, inLeft, rootIndex-1)
		root.Right = build(preLeft+leftSize+1, preRight, rootIndex+1, inRight)
		return root
	}
	return build(0, len(preorder)-1, 0, len(preorder)-1)
}

// 思路总结：以前序遍历定根，以中序遍历定左右子树，整体为先序遍历

func main() {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	buildTree(preorder, inorder)
}
