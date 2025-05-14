package trie

import (
	"fmt"
	"math"
)

type NumberType uint8

const (
	PrefixNumber NumberType = 0
	ExactNumber  NumberType = 1
)

type TrieManager struct {
	trie *trie
}

func NewTrieManager() *TrieManager {
	return &TrieManager{
		trie: NewTrie(),
	}
}

func (manager *TrieManager) Add(prefix, target string, numberType NumberType) error {
	if numberType == PrefixNumber {
		if !manager.trie.add(prefix, target, math.MaxInt32) {
			return fmt.Errorf("add prefix is failed")
		}
	} else if numberType == ExactNumber {
		if !manager.trie.add(prefix, target, uint32(len(prefix))) {
			return fmt.Errorf("add prefix is failed")
		}
	} else {
		return fmt.Errorf("number type {%v} is not valid", numberType)
	}

	return nil
}

func (manager *TrieManager) Remove(prefix string) error {
	if !manager.trie.remove(prefix) {
		return fmt.Errorf("remove route is failed")
	}

	return nil
}

func (manager *TrieManager) Update(prefix string, taget string, maxLength uint32) error {
	if !manager.trie.update(prefix, taget, maxLength) {
		return fmt.Errorf("update route is failed")
	}

	return nil
}

func (manager *TrieManager) Find(prefix string) (bool, string) {
	return manager.trie.find(prefix)
}

func (manager *TrieManager) ClearAll() {
	manager.trie.clearAll()
}
