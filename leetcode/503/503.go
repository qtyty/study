package main

import "fmt"

// https://leetcode-cn.com/problems/next-greater-element-ii/
// 503. 下一个更大元素 II
// 给定一个循环数组（最后一个元素的下一个元素是数组的第一个元素），输出每个元素的下一个更大元素。数字 x 的下一个更大的元素是按数组遍历顺序，这个数字之后的第一个比它更大的数，这意味着你应该循环地搜索它的下一个更大的数。如果不存在，则输出 -1。

// 示例 1:
// 输入: [1,2,1]
// 输出: [2,-1,2]
// 解释: 第一个 1 的下一个更大的数是 2；
// 数字 2 找不到下一个更大的数；
// 第二个 1 的下一个最大的数需要循环搜索，结果也是 2。
// 注意: 输入数组的长度不会超过 10000。

func nextGreaterElements(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	stack := make([]int, 0)
	for i := 2*n - 1; i >= 0; i-- {
		num := nums[i%n]
		for len(stack) > 0 && num >= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			res[i%n] = stack[len(stack)-1]
		} else {
			res[i%n] = -1
		}
		stack = append(stack, num)
	}
	return res
}

func main() {
	nums := []int{1, 2, 1}
	fmt.Println(nextGreaterElements(nums))
}
