package main

import (
	"testing"
)

func TestSolve(t *testing.T) {
	dict := newTrie()
	dict.add("hi")
	dict.add("it")
	dict.add("hit")
	dict.add("hilt")

	chars := [rows][cols]rune{
		{'x', 'x', 't', 'l'},
		{'x', 'x', 'h', 'i'},
		{'x', 'x', 't', 'l'},
		{'x', 'x', 'x', 'x'},
	}
	modifiers := [rows][cols]modifier{
		{none, none, x2Word, none},
		{none, none, x2Letter, none},
		{none, none, none, none},
		{none, none, none, none},
	}
	b := &board{chars, modifiers}

	sols := solve(b, dict)

	expected := []solution{
		solution{"hilt", 20, []position{position{1, 2}, position{1, 3},
			position{0, 3}, position{0, 2}}},
		solution{"hit", 16, []position{position{1, 2}, position{1, 3},
			position{0, 2}}},
		solution{"it", 2, []position{position{1, 3}, position{0, 2}}},
		solution{"hi", 1, []position{position{1, 2}, position{1, 3}}},
	}

	if len(sols) != len(expected) {
		t.Errorf("Expected %d solutions; found %d", len(expected), len(sols))
	}
	for i, sol := range sols {
		if !solutionsEqual(sol, expected[i]) {
			t.Errorf("Expected %v at index %d; found %v", expected[i], i, sol)
		}
	}
}

func solutionsEqual(sol1 solution, sol2 solution) bool {
	if sol1.word != sol2.word {
		return false
	}
	if sol1.score != sol2.score {
		return false
	}
	if len(sol1.path) != len(sol2.path) {
		return false
	}
	for i, pos := range sol1.path {
		if pos != sol2.path[i] {
			return false
		}
	}
	return true
}
