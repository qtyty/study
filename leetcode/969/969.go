package main

import "fmt"

func pancakeSort(arr []int) []int {
	n := len(arr)
	var res []int
	for i := n; i > 0; i-- {
		index := 0
		for k, v := range arr[:i] {
			if v > arr[index] {
				index = k
			}
		}
		if index == i-1 {
			continue
		}
		reverse(arr, 0, index)
		reverse(arr, 0, i-1)
		res = append(res, index+1, i)
	}
	return res
}

func reverse(arr []int, start, end int) {
	for i := start; i < (end-start+1)/2; i++ {
		arr[i], arr[end-i] = arr[end-i], arr[i]
	}
}

func main() {
	arr := []int{3, 2, 4, 1}
	res := pancakeSort(arr)
	fmt.Println(res)
}
