package main

import (
	"bufio"
	"fmt"
	"os"
)

// TrieNode represents a node in the Trie
type TrieNode struct {
	children map[byte]*TrieNode
	maxPop   int
}

func newTrieNode() *TrieNode {
	return &TrieNode{children: make(map[byte]*TrieNode)}
}

// Trie represents the Trie data structure
type Trie struct {
	root *TrieNode
}

func newTrie() *Trie {
	return &Trie{root: newTrieNode()}
}

// Insert a word into the Trie with its popularity
func (trie *Trie) Insert(word string, popularity int) {
	node := trie.root
	for i := 0; i < len(word); i++ {
		char := word[i]
		if _, exists := node.children[char]; !exists {
			node.children[char] = newTrieNode()
		}
		node = node.children[char]
		if popularity > node.maxPop {
			node.maxPop = popularity
		}
	}
}

// Autocomplete finds the maximum popularity for a given prefix
func (trie *Trie) Autocomplete(prefix string) int {
	node := trie.root
	for i := 0; i < len(prefix); i++ {
		char := prefix[i]
		if _, exists := node.children[char]; !exists {
			return -1
		}
		node = node.children[char]
	}
	return node.maxPop
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var N, Q int
	fmt.Fscan(reader, &N, &Q)

	trie := newTrie()

	// Read dictionary and insert into the Trie
	for i := 0; i < N; i++ {
		var word string
		var popularity int
		fmt.Fscan(reader, &word, &popularity)
		trie.Insert(word, popularity)
	}

	// Process queries
	currentString := ""
	for i := 0; i < Q; i++ {
		query, _ := reader.ReadString('\n')
		if query[0] == '+' {
			currentString += string(query[2])
		} else if query[0] == '-' {
			if len(currentString) > 0 {
				currentString = currentString[:len(currentString)-1]
			}
		}
		fmt.Println(trie.Autocomplete(currentString))
	}
}
