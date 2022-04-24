package main

import (
	"fmt"
	"math"
)

func binaryGap(n int) int {
	start, end := 0, math.MaxInt64
	a := n
	ans := 0
	for a > 0 {
		l := a % 2
		if l == 1 {
			if start-end > ans {
				ans = start - end
			}
			end = start
		}
		a /= 2
		start++
	}
	return ans
}

func main() {
	n := 8
	fmt.Println(binaryGap(n))
}
