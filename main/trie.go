package main

type trie struct {
	root *node
}

func newTrie() *trie {
	return &trie{&node{}}
}

func (t trie) add(str string) bool {
	if !validString(str) {
		return false
	}
	curNode := t.root
	for _, char := range str {
		nextNode := curNode.getChild(char)
		if nextNode == nil {
			nextNode = &node{}
			curNode.setChild(char, nextNode)
		}
		curNode = nextNode
	}
	curNode.wordEnd = true
	return true
}

func (t trie) contains(str string) bool {
	curNode := t.root
	for _, char := range str {
		nextNode := curNode.getChild(char)
		if nextNode == nil {
			return false
		}
		curNode = nextNode
	}
	return curNode.wordEnd
}

type node struct {
	wordEnd  bool
	children [26]*node
}

func (n *node) isWordEnd() bool {
	return n.wordEnd
}

func (n *node) getChild(char rune) *node {
	return n.children[translate(char)]
}

func (n *node) setChild(char rune, child *node) {
	n.children[translate(char)] = child
}
