package main

import "fmt"

// 标准解法：二分递归
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	total := len(nums1) + len(nums2)
	if total%2 == 1 {
		return float64(getKthElement(nums1, nums2, total/2+1))
	}
	left := float64(getKthElement(nums1, nums2, total/2))
	right := float64(getKthElement(nums1, nums2, total/2+1))
	return (left + right) / 2.0
}

func getKthElement(nums1, nums2 []int, k int) int {
	if len(nums1) > len(nums2) {
		return getKthElement(nums2, nums1, k)
	}
	if len(nums1) == 0 {
		return nums2[k-1]
	}
	if k == 1 {
		return min(nums1[0], nums2[0])
	}
	i := min(k/2, len(nums1)) - 1
	j := min(k/2, len(nums2)) - 1
	if nums1[i] <= nums2[j] {
		return getKthElement(nums1[i+1:], nums2, k-(i+1))
	} else {
		return getKthElement(nums1, nums2[j+1:], k-(j+1))
	}
}

// 思路总结：这道题的难点在与如何从两个数组中查询第 k 小的元素

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}
