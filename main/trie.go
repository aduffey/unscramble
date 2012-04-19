package main

import (
	"strings"
)

const validChars = "abcdefghijklmnopqrstuvwxyz"

type trie struct {
	root *node
}

func newTrie() *trie {
	return &trie{&node{}}
}

func (t trie) add(str string) bool {
	if !valid(str) {
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

func valid(str string) bool {
	for _, char := range str {
		if !strings.ContainsRune(validChars, char) {
			return false
		}
	}
	return true
}

func translate(char rune) int {
	return strings.IndexRune(validChars, char)
}
