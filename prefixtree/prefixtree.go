package main

import (
	"fmt")

// 前缀树/字典树：利用字符串的公共前缀来减少查询时间，最大限度
// 的减少无谓的字符串比较，查询效率比哈希树高

// 特征：1.根节点不包括字符
// 2.每个节点的所有子节点包括的字符都不相同
// 3.一个节点的所有子节点都有相同的前缀

type Trie struct {
	root *Node
}

func NewTrie() *Trie {
	return &Trie{root: NewNode()}
}

type Node struct {
	isend bool
	Children map[string]*Node // 不用slice不用遍历
}

func NewNode() *Node {
	return &Node{Children: make(map[string]*Node)}
}

func (this *Trie) Insert(s string) {
	current := this.root
	for _,item := range []rune(s) {
		if _,ok := current.Children[string(item)]; !ok {
			current.Children[string(item)] = NewNode()
		}

		current = current.Children[string(item)]
	}

	current.isend = true
}

func (this *Trie) Search(s string) bool {
	current := this.root
	for _, item := range []rune(s) {
		if _,ok := current.Children[string(item)]; !ok {
			return false
		}

		current = current.Children[string(item)]
	}

	return current.isend
}

func test() {
	strs := []string{"go","gin","golang","goapp","guest"}
	tree := NewTrie()
	for _,s := range strs {
		tree.Insert(s)
	}

	strs = append(strs,"abc","gogo","g")
	for _,s := range strs {
		fmt.Println(tree.Search(s))
	}
}

func main()  {
	test()
}