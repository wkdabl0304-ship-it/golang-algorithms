package main

import "fmt"

// 标准解法：树的数据结构建模

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func Constructor0208() Trie {
	return Trie{}
}

func (t *Trie) Insert(word string) {
	node := t
	for _, v := range word {
		index := v - 'a'
		if node.children[index] == nil {
			node.children[index] = &Trie{}
		}
		node = node.children[index]
	}
	node.isEnd = true
}

func (t *Trie) Search(word string) bool {
	node := t
	for _, v := range word {
		index := v - 'a'
		if node.children[index] == nil {
			return false
		}
		node = node.children[index]
	}
	return node.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	node := t
	for _, v := range prefix {
		index := v - 'a'
		if node.children[index] == nil {
			return false
		}
		node = node.children[index]
	}
	return true
}

// 思路总结
// 1.考虑到前缀搜索和单词搜索要复用数据结构，所以应该用树
// 2.考虑到下一个字母只有可能是 26 个字母中的一个，所以数据结构有 26 个子树的指针
// 3.对于如何区分前缀搜索和单词搜索，需要在结点引入 isEnd 字段作为判断依据

func main() {
	obj := Constructor0208()
	obj.Insert("word")
	fmt.Println(obj.Search("word"))
	fmt.Println(obj.StartsWith("wo"))
}
