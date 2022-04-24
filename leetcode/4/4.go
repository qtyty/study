package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	total := len(nums1) + len(nums2)

	if total%2 == 0 {
		// 偶数长度，取中间两个数
		left := findKthNumber(nums1, 0, nums2, 0, total/2)
		right := findKthNumber(nums1, 0, nums2, 0, total/2+1)

		return float64(left+right) / 2
	} else {
		// 奇数长度，取中数
		return float64(findKthNumber(nums1, 0, nums2, 0, total/2+1))
	}
}

// 从nums1[i:] 和 nums2[j:]之间找到第k大数
func findKthNumber(nums1 []int, i int, nums2 []int, j int, k int) int {
	// 提前预处理，默认nums1待处理长度小于nums2
	if len(nums1)-i > len(nums2)-j {
		return findKthNumber(nums2, j, nums1, i, k)
	}
	// 边界场景1
	if len(nums1) == i {
		return nums2[j+k-1]
	}
	// 边界场景2
	if k == 1 {
		return min(nums1[i], nums2[j])
	}

	si := min(i+k/2, len(nums1))
	sj := j + k/2
	// 二分
	if nums1[si-1] > nums2[sj-1] {
		return findKthNumber(nums1, i, nums2, sj, k-k/2)
	} else {
		return findKthNumber(nums1, si, nums2, j, k-(si-i))
	}

}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	nums1, nums2 := []int{1, 2}, []int{3, 4}
	median := findMedianSortedArrays(nums1, nums2)
	fmt.Println(median)
}
