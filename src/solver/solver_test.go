package solver

import (
	"testing"
)

func TestScore(t *testing.T) {
	dict := NewDict()
	dict.Add("warmest")

	chars := [Rows][Cols]rune{
		{'x', 'x', 'w', 'a'},
		{'x', 'x', 'm', 'r'},
		{'x', 'x', 'e', 's'},
		{'x', 'x', 'x', 't'},
	}
	modifiers := [Rows][Cols]Modifier{
		{None, None, X2Word, None},
		{None, None, X2Letter, None},
		{None, None, None, None},
		{None, None, None, None},
	}
	b := &Board{chars, modifiers}

	s := newSolution([]Position{Position{0, 2}, Position{0, 3}, Position{1, 3},
		Position{1, 2}, Position{2, 2}, Position{2, 3}, Position{3, 3}}, b)

	if s.Word != "warmest" {
		t.Errorf("Expected \"warmest\", got \"%s\"", s.Word)
	}
	if s.Score != 44 {
		t.Errorf("Expected score 44, got %d", s.Score)
	}
}

func TestSolve(t *testing.T) {
	dict := NewDict()
	dict.Add("hi")
	dict.Add("it")
	dict.Add("hit")
	dict.Add("hilt")

	chars := [Rows][Cols]rune{
		{'x', 'x', 't', 'l'},
		{'x', 'x', 'h', 'i'},
		{'x', 'x', 't', 'l'},
		{'x', 'x', 'x', 'x'},
	}
	modifiers := [Rows][Cols]Modifier{
		{None, None, X2Word, None},
		{None, None, X2Letter, None},
		{None, None, None, None},
		{None, None, None, None},
	}
	b := &Board{chars, modifiers}

	sols := Solve(b, dict)

	expected := []*Solution{
		&Solution{"hilt", 20, []Position{Position{1, 2}, Position{1, 3},
			Position{0, 3}, Position{0, 2}}},
		&Solution{"hit", 16, []Position{Position{1, 2}, Position{1, 3},
			Position{0, 2}}},
		&Solution{"it", 2, []Position{Position{1, 3}, Position{0, 2}}},
		&Solution{"hi", 1, []Position{Position{1, 2}, Position{1, 3}}},
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

func solutionsEqual(sol1 *Solution, sol2 *Solution) bool {
	if sol1.Word != sol2.Word {
		return false
	}
	if sol1.Score != sol2.Score {
		return false
	}
	if len(sol1.Path) != len(sol2.Path) {
		return false
	}
	for i, pos := range sol1.Path {
		if pos != sol2.Path[i] {
			return false
		}
	}
	return true
}
