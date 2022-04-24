package main

import "strconv"

//寻找重复的子树

//给定一棵二叉树 root，返回所有重复的子树。
//
//对于同一类的重复子树，你只需要返回其中任意一棵的根结点即可。
//
//如果两棵树具有相同的结构和相同的结点值，则它们是重复的。

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findDuplicateSubtrees(root *TreeNode) (res []*TreeNode) {
	count := make(map[string]int)
	strmap := make(map[string]*TreeNode)

	var dfs func(root *TreeNode) string
	dfs = func(root *TreeNode) string {
		if root == nil {
			return "#"
		}
		str := strconv.Itoa(root.Val) + "," + dfs(root.Left) + "," + dfs(root.Right)
		count[str]++
		strmap[str] = root
		if count[str] == 2 {
			res = append(res, root)
		}
		return str
	}
	dfs(root)
	return res
}

func main() {

}
