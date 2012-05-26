package solver

import (
	"testing"
)

func TestDict(t *testing.T) {
	validMsg := "Couldn't add valid word \"%s\""
	invalidMsg := "Added invalid word \"%s\""
	containsValidMsg := "Trie should contain \"%s\""

	testDict := NewDict()

	// Add some valid words
	longSphinx := "sphinxofblackquartzjudgemyvow"
	if !testDict.Add(longSphinx) {
		t.Errorf(validMsg, longSphinx)
	}
	aardvark := "aardvark"
	if !testDict.Add(aardvark) {
		t.Errorf(validMsg, aardvark)
	}
	shortSphinx := "sphinx"
	if !testDict.Add(shortSphinx) {
		t.Errorf(validMsg, shortSphinx)
	}

	// Try an invalid word
	badCarrots := "Carrots"
	if testDict.Add(badCarrots) {
		t.Errorf(invalidMsg, badCarrots)
	}

	// Check the words we've added
	if !testDict.Contains(longSphinx) {
		t.Errorf(containsValidMsg, longSphinx)
	}
	if !testDict.Contains(aardvark) {
		t.Errorf(containsValidMsg, aardvark)
	}
	if !testDict.Contains(shortSphinx) {
		t.Errorf(containsValidMsg, shortSphinx)
	}
}
