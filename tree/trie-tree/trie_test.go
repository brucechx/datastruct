package trie

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

var trieTree *Trie

func initTrieTree(){
	trieTree = NewTrie()
	trieTree.Insert("go")
	trieTree.Insert("golang")
	trieTree.Insert("apple")
	//trieTree.root.traversal(0)
}

func TestTrie_Search(t *testing.T) {
	initTrieTree()
	cases := []struct{
		input string
		output bool
	}{
		{"go", true},
		{"apple", true},
		{"gol", false},
		{"app", false},
	}
	for _, cas := range cases{
		res := trieTree.Search(cas.input)
		assert.Equal(t, cas.output, res)
	}
}

func TestTrie_StartsWith(t *testing.T) {
	initTrieTree()
	cases := []struct{
		input string
		output bool
	}{
		{"go", true},
		{"apple", true},
		{"gol", true},
		{"app", true},
		{"ax", false},
	}
	for _, cas := range cases{
		res := trieTree.StartsWith(cas.input)
		assert.Equal(t, cas.output, res)
	}
}


