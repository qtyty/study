package main

// https://leetcode-cn.com/problems/satisfiability-of-equality-equations/
// 990. 等式方程的可满足性

//给定一个由表示变量之间关系的字符串方程组成的数组，每个字符串方程 equations[i] 的长度为 4，并采用两种不同的形式之一："a==b" 或 "a!=b"。在这里，a 和 b 是小写字母（不一定不同），表示单字母变量名。
//
//只有当可以将整数分配给变量名，以便满足所有给定的方程时才返回 true，否则返回 false。

//输入：["a==b","b!=a"]
//输出：false
//解释：如果我们指定，a = 1 且 b = 1，那么可以满足第一个方程，但无法满足第二个方程。没有办法分配变量同时满足这两个方程。

type UnionFind struct {
	parent []int
	count  int
}

func (uf *UnionFind) Union(p, q int) {
	pRoot := uf.Find(p)
	qRoot := uf.Find(q)
	if pRoot == qRoot {
		return
	}
	uf.parent[pRoot] = qRoot
	uf.count--
}

func (uf *UnionFind) Find(a int) int {
	if uf.parent[a] == a {
		return a
	}
	uf.parent[a] = uf.Find(uf.parent[a])
	return uf.parent[a]
}

func equationsPossible(equations []string) bool {
	uf := &UnionFind{
		parent: make([]int, 26),
		count:  26,
	}
	for i := 0; i < 26; i++ {
		uf.parent[i] = i
	}
	for _, equation := range equations {
		if equation[1] == '=' {
			uf.Union(int(equation[0]-'a'), int(equation[3]-'a'))
		}
	}
	for _, equation := range equations {
		if equation[1] == '!' {
			if uf.Find(int(equation[0]-'a')) == uf.Find(int(equation[3]-'a')) {
				return false
			}
		}
	}
	return true
}

func main() {

}
