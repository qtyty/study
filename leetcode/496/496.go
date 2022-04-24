package main

// https://leetcode-cn.com/problems/next-greater-element-i/
// 496. 下一个更大元素 I
// 给定两个没有重复元素的数组 nums1 和 nums2 ，其中nums1 是 nums2 的子集。找到 nums1 中每个元素在 nums2 中的下一个比其大的值。

// nums1 中数字 x 的下一个更大元素是指 x 在 nums2 中对应位置的右边的第一个比 x 大的元素。如果不存在，对应位置输出-1。
// 示例 1:
//输入：nums1 = [4,1,2], nums2 = [1,3,4,2].
//输出：[-1,3,-1]
//解释：nums1 中每个值的下一个更大元素如下所述：
//- 4 ，用加粗斜体标识，nums2 = [1,3,4,2]。不存在下一个更大元素，所以答案是 -1 。
//- 1 ，用加粗斜体标识，nums2 = [1,3,4,2]。下一个更大元素是 3 。
//- 2 ，用加粗斜体标识，nums2 = [1,3,4,2]。不存在下一个更大元素，所以答案是 -1 。

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	stack := make([]int, 0)
	for i := len(nums2) - 1; i >= 0; i-- {
		n := nums2[i]
		for len(stack) > 0 && n >= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			m[n] = stack[len(stack)-1]
		} else {
			m[n] = -1
		}
		stack = append(stack, n)
	}
	res := make([]int, len(nums1))
	for i, v := range nums1 {
		res[i] = m[v]
	}
	return res
}

func main() {

}
