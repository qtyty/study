package main

import (
	"container/heap"
	"fmt"
	"sort"
)

//优势洗牌
//给定两个大小相等的数组 A 和 B，A 相对于 B 的优势可以用满足 A[i] > B[i] 的索引 i 的数目来描述。
//
//返回 A 的任意排列，使其相对于 B 的优势最大化。

//输入：A = [2,7,11,15], B = [1,10,4,11]
//输出：[2,11,7,15]

type hp [][]int

func (h hp) Len() int {
	return len(h)
}

func (h hp) Less(i, j int) bool {
	return h[i][0] > h[j][0]
}

func (h hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *hp) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *hp) Pop() interface{} {
	old := *h
	n := len(old)
	a := old[n-1]
	*h = old[:n-1]
	return a
}

func advantageCount(nums1 []int, nums2 []int) []int {
	h := &hp{}
	heap.Init(h)

	for i, num := range nums2 {
		heap.Push(h, []int{num, i})
	}

	sort.Ints(nums1)

	n := len(nums1)
	ans := make([]int, n)
	left, right := 0, n-1
	for h.Len() > 0 {
		if x := heap.Pop(h).([]int); x[0] < nums1[right] {
			ans[x[1]] = nums1[right]
			right--
		} else {
			ans[x[1]] = nums1[left]
			left++
		}
	}
	return ans
}

func main() {
	nums1 := []int{2, 7, 11, 15}
	nums2 := []int{1, 10, 4, 11}
	fmt.Println(advantageCount(nums1, nums2))
}
