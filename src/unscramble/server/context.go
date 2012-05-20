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

// The context for the template.
type context struct {
	Board     [solver.Rows][solver.Cols]*cell
	Solutions []*solver.Solution
}

func newContext(b *solver.Board, sols []*solver.Solution) *context {
	ctxBoard := [solver.Rows][solver.Cols]*cell{}
	for i, row := range b.Chars {
		for j, v := range row {
			var value string
			if v == 'q' {
				value = "Qu"
			} else {
				value = strings.ToUpper(string(v))
			}

			var modifier string
			m := b.Modifiers[i][j]
			if m == solver.X2Letter {
				modifier = x2LetterClass
			} else if m == solver.X2Word {
				modifier = x2WordClass
			} else if m == solver.X3Letter {
				modifier = x3LetterClass
			} else if m == solver.X3Word {
				modifier = x3WordClass
			} else {
				modifier = ""
			}

			ctxBoard[i][j] = &cell{value, modifier}
		}
	}
	return &context{ctxBoard, sols}
}

func emptyContext() *context {
	ctxBoard := [solver.Rows][solver.Cols]*cell{}
	for i, row := range ctxBoard {
		for j, _ := range row {
			ctxBoard[i][j] = &cell{"", ""}
		}
	}
	sols := []*solver.Solution{}
	return &context{ctxBoard, sols}
}

type cell struct {
	Value string
	Class string
}
