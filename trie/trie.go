package trie

type charNode struct {
	parent *charNode
	depth  uint32
	index  int32

	hasValueFlag bool
	target       string
	maxLength    uint32

	numberOfChildren uint8
	children         [10]*charNode
}

func newCharNode(parent *charNode, depth uint32, index int32) *charNode {
	n := &charNode{
		parent:           parent,
		depth:            depth,
		index:            index,
		hasValueFlag:     false,
		maxLength:        0,
		numberOfChildren: 0,
	}
	return n
}

func (n *charNode) clean() {
	for i := range len(n.children) {
		if n.children[i] != nil {
			n.children[i].clean()
			n.children[i] = nil
		}
	}

	n.numberOfChildren = 0
}

type trie struct {
	tree_root *charNode
}

func NewTrie() *trie {
	return &trie{
		tree_root: newCharNode(nil, 0, -1),
	}
}

func (t *trie) add(prefix string, target string, maxLength uint32) bool {
	n := t.createSubTree(prefix, 0, uint32(len(prefix)), t.tree_root)
	if n != nil {
		if !n.hasValueFlag {
			n.target = target
			n.maxLength = maxLength
			n.hasValueFlag = true
			return true
		}
	}

	return false
}

func (t *trie) remove(prefix string) bool {
	n := t.findMatching(prefix, 0, uint32(len(prefix)), t.tree_root, nil)
	if (n != nil) && (n.depth == uint32(len(prefix))) {
		n.hasValueFlag = false
		t.cleanUp(n)
		return true
	}

	return false
}

func (t *trie) update(prefix string, target string, maxLength uint32) bool {
	n := t.findMatching(prefix, 0, uint32(len(prefix)), t.tree_root, nil)
	if (n != nil) && (n.depth == uint32(len(prefix))) {
		n.target = target
		n.maxLength = maxLength
		n.hasValueFlag = true

		return true
	}

	return false
}

func (t *trie) clearAll() {
	t.tree_root.clean()
}

func (t *trie) find(prefix string) (bool, string) {
	n := t.findMatching(prefix, 0, uint32(len(prefix)), t.tree_root, nil)
	if n != nil {
		return true, n.target
	}

	return false, ""
}

func (t *trie) cleanUp(current *charNode) {
	if (!current.hasValueFlag) && (current.numberOfChildren == 0) {
		parent := current.parent
		if parent != nil {
			parent.numberOfChildren--
			parent.children[current.index] = nil
			t.cleanUp(parent)
		}
	}
}

func (t *trie) createSubTree(prefix string, first uint32, last uint32, current *charNode) *charNode {
	if first < last {
		index := prefix[first] - '0'
        if index > 9 {
            return nil
        }

		if current.children[index] == nil {
			current.children[index] = newCharNode(current, current.depth+1, int32(index))
			current.numberOfChildren++
		}

		return t.createSubTree(prefix, first+1, last, current.children[index])
	}

	if current.hasValueFlag {
		return nil
	}

	return current
}

func (t *trie) findMatching(prefix string, first uint32, last uint32, current *charNode, lastValidValue *charNode) *charNode {
	if current.hasValueFlag &&
		(current.maxLength >= uint32(len(prefix))) {
		lastValidValue = current
	}

	if first < last {
		index := prefix[first] - '0'
        if index > 9 {
            return nil
        }

		if current.children[index] != nil {
			return t.findMatching(prefix, first+1, last, current.children[index], lastValidValue)
		}
	}

	return lastValidValue
}
