package solver

type Dict struct {
	root *node
}

func NewDict() *Dict {
	return &Dict{&node{}}
}

func (d Dict) Add(str string) bool {
	if !ValidString(str) {
		return false
	}
	curNode := d.root
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

func (d Dict) Contains(str string) bool {
	curNode := d.root
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
