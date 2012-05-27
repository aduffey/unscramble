package solver

// Dict represents a dictionary that can store arbitrary strings. It is
// essentially a set of strings, not storing duplicates, that permits insertion
// and membership testing.
type Dict struct {
	// This root node is special. It represents the empty string. Its value
	// field is ignored, and it will never have a "next" node (only children).
	// This also allows us to avoid treating the empty dictionary as a special
	// case in our methods.
	root *node
}

// NewDict creates an empty dictionary.
func NewDict() *Dict {
	return &Dict{&node{}}
}

// Add inserts a string into the dictionary. Returns true if the string was
// successfully added or false if the string was already in the dictionary.
func (d Dict) Add(str string) bool {
	curNode := d.root
	for i := 0; i < len(str); i++ {
		curNode = curNode.addChild(str[i])
	}
	existed := curNode.wordEnd
	curNode.wordEnd = true
	return !existed
}

// Contains returns true if and only if the dictionary contains the given
// string.
func (d Dict) Contains(str string) bool {
	curNode := d.root
	for i := 0; i < len(str); i++ {
		curNode = curNode.getChild(str[i])
		if curNode == nil {
			return false
		}
	}
	return curNode.wordEnd
}

// node represents a node in the trie. The children of each node are stored as a
// linked list with each node holding a pointer to its sibling in addition to a
// pointer to its first child. This significantly reduces memory usage compared
// to storing a dense array of children with one element for each possible child
// pointer, most of which would be nil in a large trie.
type node struct {
	value byte
	wordEnd  bool
	next *node
	child *node
}

// getChild retrieves the child of the current node for the given value. Returns
// nil if no such child exists.
func (n *node) getChild(b byte) *node {
	child := n.child
	for child != nil && child.value != b {
		child = child.next
	}
	return child
}

// addChild adds a child node with the given value to this node and returns the
// new child. If a child with the given value already exists, it is returned
// instead and no new node is added.
func (n *node) addChild(b byte) *node {
	child := (*node)(nil)
	if n.child == nil {
		// If this node has no child, we simply add one
		n.child = &node{b, false, nil, nil}
		child = n.child
	} else {
		// Otherwise we have to iterate through all the children
		c := n.child
		for {
			if c.value == b {
				// If we see a child with the value we're going to add, return
				// the child instead
				child = c
				break
			} else if c.next == nil {
				// If we hit the end of the child list, insert a new child
				c.next = &node{b, false, nil, nil}
				child = c.next
				break
			}
			c = c.next
		}
	}
	return child
}
