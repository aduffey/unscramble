package main

import (
	"strings"
	"unscramble/solver"
)

// The CSS class names corresponding to each modifier. These must agree with the
// classes declared in the Javascript and CSS files!
const (
	x2LetterClass = "x2Letter"
	x2WordClass   = "x2Word"
	x3LetterClass = "x3Letter"
	x3WordClass   = "x3Word"
)

// The CSS class names corresponding to a cell that is part of a specific
// solution. Like above these must agree with classes declared elsewhere!
const (
	firstLetterClass = "firstLetter"
	inSolutionClass  = "inSolution"
)

// The context for the template.
type context struct {
	Board     [solver.Rows][solver.Cols]*cell
	Solutions []*solution
}

func newContext(b *solver.Board, sols []*solver.Solution) *context {
	// Make the main board
	ctxBoard := populateBoardLetters(b)
	// Set the modifiers
	for i, row := range b.Modifiers {
		for j, mod := range row {
			class := ""
			if mod == solver.X2Letter {
				class = x2LetterClass
			} else if mod == solver.X2Word {
				class = x2WordClass
			} else if mod == solver.X3Letter {
				class = x3LetterClass
			} else if mod == solver.X3Word {
				class = x3WordClass
			}
			ctxBoard[i][j].Class = class
		}
	}

	// Make the solutions
	ctxSols := make([]*solution, len(sols))
	for i, sol := range sols {
		solBoard := populateBoardLetters(b)
		for i, pos := range sol.Path {
			if i == 0 {
				solBoard[pos.Row][pos.Col].Class = firstLetterClass
			} else {
				solBoard[pos.Row][pos.Col].Class = inSolutionClass
			}
		}
		ctxSols[i] = &solution{sol.Word, sol.Score, solBoard}
	}
	return &context{ctxBoard, ctxSols}
}

func emptyContext() *context {
	ctxBoard := [solver.Rows][solver.Cols]*cell{}
	for i, row := range ctxBoard {
		for j, _ := range row {
			ctxBoard[i][j] = &cell{"", ""}
		}
	}
	sols := []*solution{}
	return &context{ctxBoard, sols}
}

type cell struct {
	Value string
	Class string
}

type solution struct {
	Word  string
	Score int
	Board [solver.Rows][solver.Cols]*cell
}

func populateBoardLetters(b *solver.Board) [solver.Rows][solver.Cols]*cell {
	ctxBoard := [solver.Rows][solver.Cols]*cell{}
	for i, row := range b.Chars {
		for j, v := range row {
			var value string
			if v == 'q' {
				value = "Qu"
			} else {
				value = strings.ToUpper(string(v))
			}
			ctxBoard[i][j] = &cell{value, ""}
		}
	}
	return ctxBoard
}
