package trie

// Node represents Trie node
type Node struct {
	// map used to quick check if character is already inserted and fetching accompanied node
	children map[rune]*Node
	word     bool
}

// Trie represents trie data structure
type Trie struct {
	root *Node
}

// Constructor Initialize your data structure here
func Constructor() Trie {
	return Trie{
		root: &Node{children: make(map[rune]*Node)},
	}
}

// Insert Inserts a word into the trie
func (t *Trie) Insert(word string) {
	currentNode := t.root
	for _, r := range word {
		if _, ok := currentNode.children[r]; !ok {
			currentNode.children[r] = &Node{children: make(map[rune]*Node)}
		}
		currentNode = currentNode.children[r]
	}
	currentNode.word = true
}

// find returns the Node of character in s
func (t *Trie) find(s string) *Node {
	currentNode := t.root
	for _, r := range s {
		if _, ok := currentNode.children[r]; !ok {
			return nil
		}
		currentNode = currentNode.children[r]
	}
	return currentNode
}

// Search returns if the word is in the trie
func (t *Trie) Search(word string) bool {
	foundNode := t.find(word)
	return foundNode != nil && foundNode.word
}

// StartsWith returns if there is any word in the trie that starts with the given prefix
func (t *Trie) StartsWith(prefix string) bool {
	return t.find(prefix) != nil
}
