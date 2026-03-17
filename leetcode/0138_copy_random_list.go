package main

import "fmt"

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// 标准解法：哈希法，还有原地拆分法
func copyRandomList(head *Node) *Node {
	// 1.初始化 map
	m := make(map[*Node]*Node)
	for a := head; a != nil; a = a.Next {
		m[a] = &Node{Val: a.Val}
	}
	// 2.连线
	for p := head; p != nil; p = p.Next {
		m[p].Next = m[p.Next]
		m[p].Random = m[p.Random]
	}
	return m[head]
}

func main() {
	a1, a2, a3, a4, a5 := new(Node), new(Node), new(Node), new(Node), new(Node)
	a1.Val, a2.Val, a3.Val, a4.Val, a5.Val = 7, 13, 11, 10, 1
	a1.Next, a2.Next, a3.Next, a4.Next = a2, a3, a4, a5
	a1.Random, a2.Random, a3.Random, a4.Random, a5.Random = nil, a1, a5, a3, a1
	for p := copyRandomList(a1); p != nil; p = p.Next {
		fmt.Print(p.Val, " ")
	}
}
