package main

import (
	"log"
	"go-trie-manager/trie"
)

var TrieManagerInstance *trie.TrieManager

func find(number string) {
	res, user := TrieManagerInstance.Find(number)
	if !res {
		log.Printf("could not find number {%s}\n", number)
	} else {
		log.Printf("find user {%s} for number {%s}\n", user, number)
	}
}

func add(prefix, target string, numberType trie.NumberType) {
	err := TrieManagerInstance.Add(prefix, target, numberType)
	if err != nil {
		log.Printf("add prefix {%s} is failed\n", prefix)
	} else {
		log.Printf("successfully add prefix {%s}\n", prefix)
	}
}

func remove(prefix string) {
	err := TrieManagerInstance.Remove(prefix)
	if err != nil {
		log.Printf("remove prefix {%s} is failed\n", prefix)
	} else {
		log.Printf("successfully remove prefix {%s}\n", prefix)
	}
}

func main() {
	TrieManagerInstance = trie.NewTrieManager()

	add("982000", "majid", trie.PrefixNumber)
	add("982000456", "mahdi", trie.ExactNumber)
	add("9830007863", "ali", trie.ExactNumber)
	add("9830009856", "saeed", trie.PrefixNumber)

	add("982000", "meysam", trie.PrefixNumber)

    add("982000", "mohammad", trie.PrefixNumber)

    add("a123", "wrong", trie.PrefixNumber)

	find("982000123")
	find("982000456")
	find("9820004567")
	find("9820004467")

	find("9830007863")
	find("9830009856")
	find("98300098569")
	find("983000198569")

	remove("982000456")
    remove("982000456")

	find("982000456")

    find("test")
}
