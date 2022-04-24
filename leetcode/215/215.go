package main

import (
	"fmt"
	"math/rand"
)

//find the Kth largest element in an array.
func findKthLargest(nums []int, k int) int {
	return quickSelect(nums, k, 0, len(nums)-1)
}

func partition(nums []int, left, right int) int {
	pivot := rand.Intn(right-left+1) + left
	nums[pivot], nums[right] = nums[right], nums[pivot]
	i := left
	for j := left; j < right; j++ {
		if nums[j] >= nums[right] {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[i], nums[right] = nums[right], nums[i]
	return i
}

func quickSelect(nums []int, k int, left, right int) int {
	pivot := partition(nums, left, right)
	if pivot == k-1 {
		return nums[pivot]
	} else if pivot > k-1 {
		return quickSelect(nums, k, left, pivot-1)
	} else {
		return quickSelect(nums, k, pivot+1, right)
	}
}

func main() {
	nums := []int{3, 2, 1, 5, 6, 4}
	fmt.Println(findKthLargest(nums, 6))

}
