package solver

import (
	"bytes"
	"fmt"
	"sort"
)

// ----- Describe the board -----

const (
	// The size of the board
	Rows = 4
	Cols = 4
)

// A score modifier applied to position on the board.
type Modifier int

const (
	None     Modifier = iota
	X2Letter Modifier = iota
	X2Word   Modifier = iota
	X3Letter Modifier = iota
	X3Word   Modifier = iota
)

var letterVals = map[rune]int{
	'a': 1,
	'b': 4,
	'c': 4,
	'd': 2,
	'e': 1,
	'f': 1,
	'g': 3,
	'h': 3,
	'i': 1,
	'j': 10,
	'k': 5,
	'l': 2,
	'm': 4,
	'n': 2,
	'o': 1,
	'p': 4,
	'q': 10, // Really 'qu'; 'q' does not appear alone
	'r': 1,
	's': 1,
	't': 1,
	'u': 2,
	'v': 5,
	'w': 4,
	'x': 8,
	'y': 3,
	'z': 10,
}

// Represents the game board.
type Board struct {
	Chars     [Rows][Cols]rune
	Modifiers [Rows][Cols]Modifier
}

// Create a game board from a string representation. The format is a simple list
// of letters on the board, in row major order. Prepending a letter with "2" or
// "3" specifies that cell as a double or triple letter bonus. Prepending a
// letter with "22" or "33" specifies that cell as a double or triple word
// bonus.
//
// A "Qu" cell is represented simply by a "q" character.
//
// Thus, the following string:
//
//     2abcdefghij22klmnqp
//
// represents the board:
//
//     A  B  C  D
//     E  F  G  H
//     I  J  K  L
//     M  N  Qu P
//
// where the "A" has a double letter bonus and the "K" has a double word bonus.
//
// If an error occurs while parsing, a non-nil error value will be returned.
func NewBoardFromString(boardString string) (*Board, error) {
	b := &Board{}
	row, col := 0, 0
	for index, char := range boardString {
		if row >= Rows || col >= Cols {
			err := fmt.Sprintf("Input string is too long, expected %d cells",
				Rows*Cols)
			return nil, parseError(err)
		}

		if char == '2' {
			if b.Modifiers[row][col] == None {
				b.Modifiers[row][col] = X2Letter
			} else if b.Modifiers[row][col] == X2Letter {
				b.Modifiers[row][col] = X2Word
			} else {
				err := fmt.Sprintf("Unexpected \"2\" at index %v", index)
				return nil, parseError(err)
			}
		} else if char == '3' {
			if b.Modifiers[row][col] == None {
				b.Modifiers[row][col] = X3Letter
			} else if b.Modifiers[row][col] == X3Letter {
				b.Modifiers[row][col] = X3Word
			} else {
				err := fmt.Sprintf("Unexpected \"3\" at index %d", index)
				return nil, parseError(err)
			}
		} else if ValidChar(char) {
			b.Chars[row][col] = char
			col++
			if col >= Cols {
				row++
				col = 0
			}
		} else {
			err := fmt.Sprintf("Invalid symbol %v at index %d", char, index)
			return nil, parseError(err)
		}
	}
	if row < (Rows-1) || row == (Rows-1) && col < (Cols-1) {
		err := fmt.Sprintf("Input string is too short, expected %d cells",
			Rows*Cols)
		return nil, parseError(err)
	}
	return b, nil
}

type parseError string

func (pe parseError) Error() string {
	return string(pe)
}

// Represents the position of a cell on the board.
type Position struct {
	Row int
	Col int
}

// ----- Describe the solution -----

// A solution to the game.
type Solution struct {
	Word  string
	Score int
	Path  []Position
}

func newSolution(path []Position, b *Board) *Solution {
	var buf bytes.Buffer

	score := 0
	x2WordModSet := false
	x3WordModSet := false
	for _, pos := range path {
		char := b.Chars[pos.Row][pos.Col]
		buf.WriteRune(char)

		// Special case: 'q' never occurs by itself, only as 'qu'
		if char == 'q' {
			buf.WriteRune('u')
		}

		switch b.Modifiers[pos.Row][pos.Col] {
		case X2Letter:
			score += letterVals[char] * 2
		case X3Letter:
			score += letterVals[char] * 3
		case X2Word:
			score += letterVals[char]
			x2WordModSet = true
		case X3Word:
			score += letterVals[char]
			x3WordModSet = true
		case None:
			score += letterVals[char]
		}
	}
	word := buf.String()

	// We use the length of the path for scoring, rather than the length of the
	// actual word, because 'qu' is counted as one letter for scoring purposes
	length := len(path)

	// Special case: 2-letter words are one point, and only word-modifiers
	// affect them
	if length == 2 {
		score = 1
	}

	if x2WordModSet {
		score *= 2
	}
	if x3WordModSet {
		score *= 3
	}

	// Word length bonus
	if length == 5 {
		score += 3
	} else if length == 6 {
		score += 6
	} else if length == 7 {
		score += 10
	} else if length > 7 {
		score += 5*(length-7) + 10
	}

	return &Solution{word, score, path}
}

// ----- Do the solving -----

// Holds a slice containing the neighbors for each board position. We can
// precompute this.
var neighbors [Rows][Cols][]Position

func adjacent(pos Position) []Position {
	var adj []Position
	startRow := pos.Row - 1
	if startRow < 0 {
		startRow = 0
	}
	for row := startRow; row <= pos.Row+1 && row < Rows; row++ {
		startCol := pos.Col - 1
		if startCol < 0 {
			startCol = 0
		}
		for col := startCol; col <= pos.Col+1 && col < Cols; col++ {
			adj = append(adj, Position{row, col})
		}
	}
	return adj
}

func init() {
	for row := 0; row < Rows; row++ {
		for col := 0; col < Cols; col++ {
			neighbors[row][col] = adjacent(Position{row, col})
		}
	}
}

// We just need this to sort the solutions
type solutions []*Solution

func (sols solutions) Len() int {
	return len(sols)
}

func (sols solutions) Less(i int, j int) bool {
	// So we order by *decreasing* score
	return sols[i].Score > sols[j].Score
}

func (sols solutions) Swap(i int, j int) {
	sols[j], sols[i] = sols[i], sols[j]
}

// Find all the solutions on the board. The solutions returned will be ordered
// by score, highest to lowest. If a word can be formed two or more ways, the
// solutions will only contain the highest scoring possibility.
func Solve(b *Board, dict *Dict) []*Solution {
	var sols solutions
	for i, row := range b.Chars {
		for j, _ := range row {
			solveHelper(Position{i, j}, []Position{}, b, dict.root, &sols)
		}
	}

	sort.Sort(sols)

	// Generate a new list of solutions holding only unique solutions with the
	// highest score
	uniqueSols := make([]*Solution, 0, len(sols))
	uniques := NewDict()
	for _, sol := range sols {
		if !uniques.Contains(sol.Word) {
			uniqueSols = append(uniqueSols, sol)
			uniques.Add(sol.Word)
		}
	}

	return uniqueSols
}

func solveHelper(pos Position, path []Position, b *Board, curNode *node,
	sols *solutions) {
	// Update our position in the trie
	curChar := b.Chars[pos.Row][pos.Col]
	curNode = curNode.getChild(curChar)
	// Special case: 'q' never occurs by itself, only as 'qu'
	if curChar == 'q' && curNode != nil {
		curNode = curNode.getChild('u')
	}

	// Base case: we are at a leaf in the trie, so there are no more possible
	// words along this path
	if curNode == nil {
		return
	}

	// Otherwise we are in the recursive case
	path = append(path, pos)
	if curNode.isWordEnd() {
		// We need a copy of the path, because we are mutating it as we go
		copyPath := make([]Position, len(path))
		copy(copyPath, path)
		*sols = append(*sols, newSolution(copyPath, b))
	}
	cell := path[len(path)-1]
	for _, nextPos := range neighbors[cell.Row][cell.Col] {
		// Check for cycles
		if !inPath(path, nextPos) {
			solveHelper(nextPos, path, b, curNode, sols)
		}
	}
}

func inPath(path []Position, pos Position) bool {
	for _, p := range path {
		if pos == p {
			return true
		}
	}
	return false
}
