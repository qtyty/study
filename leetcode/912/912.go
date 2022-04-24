package main

import (
	"fmt"
	"math/rand"
)

//quick sortArray
func sortArray(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func partition(nums []int, left int, right int) int {
	pivot := rand.Intn(right-left+1) + left
	//swap pivot to the end
	nums[pivot], nums[right] = nums[right], nums[pivot]
	i := left
	for j := left; j < right; j++ {
		if nums[j] < nums[right] {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[i], nums[right] = nums[right], nums[i]
	return i
}

func quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}

	pivot := partition(nums, left, right)
	quickSort(nums, left, pivot-1)
	quickSort(nums, pivot+1, right)
}

func main() {
	nums := []int{5, 2, 3, 1}
	fmt.Println(sortArray(nums))
}
