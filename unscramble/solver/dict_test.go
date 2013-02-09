package solver

import (
	"testing"
)

func TestDict(t *testing.T) {
	validMsg := "Couldn't add new word \"%s\""
	duplicateMsg := "Added duplicate word \"%s\""
	containsMsg := "Dict should contain \"%s\""

	testDict := &Dict{}

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

	// Try a duplicate word
	if testDict.Add(shortSphinx) {
		t.Errorf(duplicateMsg, shortSphinx)
	}

	// Check the words we've added
	if !testDict.Contains(longSphinx) {
		t.Errorf(containsMsg, longSphinx)
	}
	if !testDict.Contains(aardvark) {
		t.Errorf(containsMsg, aardvark)
	}
	if !testDict.Contains(shortSphinx) {
		t.Errorf(containsMsg, shortSphinx)
	}
}
