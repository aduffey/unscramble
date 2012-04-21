package main

import (
	"testing"
)

func TestTrie(t *testing.T) {
	validMsg := "Couldn't add valid word \"%s\""
	invalidMsg := "Added invalid word \"%s\""
	containsValidMsg := "Trie should contain \"%s\""

	testTrie := newTrie()

	// Add some valid words
	longSphinx := "sphinxofblackquartzjudgemyvow"
	if !testTrie.add(longSphinx) {
		t.Errorf(validMsg, longSphinx)
	}
	aardvark := "aardvark"
	if !testTrie.add(aardvark) {
		t.Errorf(validMsg, aardvark)
	}
	shortSphinx := "sphinx"
	if !testTrie.add(shortSphinx) {
		t.Errorf(validMsg, shortSphinx)
	}

	// Try an invalid word
	badCarrots := "Carrots"
	if testTrie.add(badCarrots) {
		t.Errorf(invalidMsg, badCarrots)
	}

	// Check the words we've added
	if !testTrie.contains(longSphinx) {
		t.Errorf(containsValidMsg, longSphinx)
	}
	if !testTrie.contains(aardvark) {
		t.Errorf(containsValidMsg, aardvark)
	}
	if !testTrie.contains(shortSphinx) {
		t.Errorf(containsValidMsg, shortSphinx)
	}
}
