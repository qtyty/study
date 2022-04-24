package main

import (
	"fmt"
	"math"
)

//字符的最短距离
func shortestToChar(s string, c byte) []int {
	n := len(s)
	ans := make([]int, n)

	left := math.MinInt32
	for i, ch := range s {
		if byte(ch) == c {
			left = i
		} else {
			ans[i] = i - left
		}
	}

	right := math.MaxInt32
	for i := n - 1; i >= 0; i-- {
		if s[i] == c {
			right = i
		} else {
			ans[i] = min(ans[i], right-i)
		}
	}

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	//输入：s = "loveleetcode", c = "e"
	//输出：[3,2,1,0,1,0,0,1,2,2,1,0]
	s := "loveleetcode"
	fmt.Println(shortestToChar(s, 'e'))
}
