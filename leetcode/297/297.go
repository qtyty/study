package main

import (
	"fmt"
	"strconv"
	"strings"
)

//序列化是将一个数据结构或者对象转换为连续的比特位的操作，进而可以将转换后的数据存储在一个文件或者内存中，同时也可以通过网络传输到另一个计算机环境，采取相反方式重构得到原数据。
//
//请设计一个算法来实现二叉树的序列化与反序列化。这里不限定你的序列 / 反序列化算法执行逻辑，你只需要保证一个二叉树可以被序列化为一个字符串并且将这个字符串反序列化为原始的树结构。

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

type Codec struct {
	str string
}

func Constructor() Codec {
	return Codec{""}
}

// Serializes a tree to a single string.
func (codec Codec) serialize(root *TreeNode) string {
	sb := &strings.Builder{}
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			sb.WriteString("null,")
			return
		}
		sb.WriteString(strconv.Itoa(root.Val) + ",")
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	return sb.String()
}

// Deserializes your encoded data to tree.
func (codec Codec) deserialize(data string) *TreeNode {
	list := strings.Split(data, ",")
	var dfs func(*[]string) *TreeNode
	dfs = func(list *[]string) *TreeNode {
		str := (*list)[0]
		*list = (*list)[1:]
		if str == "null" {
			return nil
		}
		val, _ := strconv.Atoi(str)
		node := &TreeNode{Val: val}
		node.Left = dfs(list)
		node.Right = dfs(list)
		return node
	}
	return dfs(&list)
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */

func main() {
	str := "1,2,null,null,3,4,null,null,5,null,null,"
	root := Constructor().deserialize(str)
	fmt.Println(root)
}
