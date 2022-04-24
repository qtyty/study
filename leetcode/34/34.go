package main

import "fmt"

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	left, right := 0, len(nums)-1
	if nums[left] > target || nums[right] < target {
		return []int{-1, -1}
	}

	start, end := -1, -1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if nums[left] != target {
		return []int{-1, -1}
	}
	start = left

	left, right = 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			left = mid + 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	end = right
	return []int{start, end}
}

func main() {
	//输入：nums = [5,7,7,8,8,10], target = 8
	//输出：[3,4]
	nums := []int{5, 7, 7, 8, 8, 10}
	fmt.Println(searchRange(nums, 8))
}
