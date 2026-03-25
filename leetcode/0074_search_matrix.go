package main

import "fmt"

// 个人解法：二次二分法
//func searchMatrix(matrix [][]int, target int) bool {
//	arrLeft, arrRight := 0, len(matrix)-1
//	for arrLeft <= arrRight {
//		mid := arrLeft + (arrRight-arrLeft)/2
//		if target == matrix[mid][0] {
//			return true
//		} else if target > matrix[mid][0] {
//			arrLeft = mid + 1
//		} else {
//			arrRight = mid - 1
//		}
//	}
//	if arrLeft > 0 {
//		arr := matrix[arrLeft-1]
//		left, right := 0, len(arr)-1
//		for left <= right {
//			mid := left + (right-left)/2
//			if target == arr[mid] {
//				return true
//			} else if target > arr[mid] {
//				left = mid + 1
//			} else {
//				right = mid - 1
//			}
//		}
//	}
//	return false
//}

// 标准解法：一次二分法
func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	left, right := 0, m*n-1
	for left <= right {
		mid := left + (right-left)>>1
		midVal := matrix[mid/n][mid%n]
		if target == midVal {
			return true
		} else if target > midVal {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}

func main() {
	matrix := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	target := 3
	fmt.Println(searchMatrix(matrix, target))
}
