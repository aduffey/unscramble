package solver

// A dictionary that can store only strings that are valid in this game (see
// ValidString).
type Dict struct {
	root *node
}

// Create a new empty dictionary.
func NewDict() *Dict {
	return &Dict{&node{}}
}

// Add a string to the dictionary. Returns false if the input string was
// invalid, true otherwise.
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

// Check if the dictionary contains the given string.
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
	children [numValidChars]*node
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
