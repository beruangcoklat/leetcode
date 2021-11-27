type Node struct {
	val    rune
	child  map[rune]*Node
	isWord bool
}

type Trie struct {
	root *Node
}

func newNode(val rune) *Node {
	return &Node{
		val:    val,
		child:  make(map[rune]*Node),
		isWord: false,
	}
}

func Constructor() Trie {
	return Trie{
		root: newNode(' '),
	}
}

func (this *Trie) Insert(word string) {
	curr := this.root
	for _, c := range word {
		if _, ok := curr.child[c]; !ok {
			curr.child[c] = newNode(c)
		}
		curr = curr.child[c]
	}
	curr.isWord = true
}

func (this *Trie) Search(word string) bool {
	curr := this.root
	for _, c := range word {
		child, ok := curr.child[c]
		if !ok {
			return false
		}
		curr = child
	}
	return curr.isWord
}

func (this *Trie) StartsWith(prefix string) bool {
	curr := this.root
	for _, c := range prefix {
		child, ok := curr.child[c]
		if !ok {
			return false
		}
		curr = child
	}
	return true
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
