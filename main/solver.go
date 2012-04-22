package main

import (
	"bytes"
	"sort"
)

// ----- Describe the board -----

const (
	rows = 4
	cols = 4
)

type modifier int

const (
	none     modifier = iota
	x2Letter modifier = iota
	x2Word   modifier = iota
	x3Letter modifier = iota
	x3Word   modifier = iota
)

var letterVals = map[rune]int{
	'a': 1,
	'b': 1,
	'c': 4,
	'd': 4,
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
	'q': 10,
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

type board struct {
	chars     [rows][cols]rune
	modifiers [rows][cols]modifier
}

type position struct {
	row int
	col int
}

// ----- Describe the solution -----

type solution struct {
	word  string
	score int
	path  []position
}

func newSolution(path []position, b *board) solution {
	var buf bytes.Buffer

	score := 0
	x2WordModSet := false
	x3WordModSet := false
	for _, pos := range path {
		char := b.chars[pos.row][pos.col]
		buf.WriteRune(char)

		switch b.modifiers[pos.row][pos.col] {
		case x2Letter:
			score += letterVals[char] * 2
		case x3Letter:
			score += letterVals[char] * 3
		case x2Word:
			score += letterVals[char]
			x2WordModSet = true
		case x3Word:
			score += letterVals[char]
			x3WordModSet = true
		case none:
			score += letterVals[char]
		}
	}
	word := buf.String()

	// Special case: 2-letter words are one point, and only word-modifiers
	// affect them
	if len(word) == 2 {
		score = 1
	}

	if x2WordModSet {
		score *= 2
	}
	if x3WordModSet {
		score *= 3
	}

	// Word length bonus
	if len(word) == 5 {
		score += 3
	} else if len(word) == 6 {
		score += 6
	} else if len(word) == 7 {
		score += 10
	} else if len(word) > 7 {
		score += 5*(len(word)-7) + 10
	}

	return solution{word, score, path}
}

// ----- Do the solving -----

type solutions []solution

func (sols solutions) Len() int {
	return len(sols)
}

func (sols solutions) Less(i int, j int) bool {
	// So we order by *decreasing* score
	return sols[i].score > sols[j].score
}

func (sols solutions) Swap(i int, j int) {
	sols[j], sols[i] = sols[i], sols[j]
}

func solve(b *board, dict *trie) []solution {
	var sols solutions
	for i, row := range b.chars {
		for j, _ := range row {
			solveHelper(position{i, j}, []position{}, b, dict.root, &sols)
		}
	}

	sort.Sort(sols)

	// Generate a new list of solutions holding only unique solutions with the
	// highest score
	uniqueSols := []solution{}
	uniques := newTrie()
	for _, sol := range sols {
		if !uniques.contains(sol.word) {
			uniqueSols = append(uniqueSols, sol)
			uniques.add(sol.word)
		}
	}

	return uniqueSols
}

func solveHelper(pos position, path []position, b *board, curNode *node,
	sols *solutions) {
	// Update to next position
	curChar := b.chars[pos.row][pos.col]
	curNode = curNode.getChild(curChar)
	path = append(path, pos)

	// Base case: we are at a leaf in the trie, so there are no more possible
	// words along this path
	if curNode == nil {
		return
	}

	// Otherwise we are in the recursive case
	if curNode.isWordEnd() {
		// We need a copy of the path, because we are mutating it as we go
		copyPath := make([]position, len(path))
		copy(copyPath, path)
		*sols = append(*sols, newSolution(copyPath, b))
	}
	for _, nextPos := range adjacent(path[len(path)-1]) {
		// Check for cycles
		if !inPath(path, nextPos) {
			solveHelper(nextPos, path, b, curNode, sols)
		}
	}
}

func adjacent(pos position) []position {
	var adj []position
	startRow := pos.row - 1
	if startRow < 0 {
		startRow = 0
	}
	for row := startRow; row <= pos.row+1 && row < rows; row++ {
		startCol := pos.col - 1
		if startCol < 0 {
			startCol = 0
		}
		for col := startCol; col <= pos.col+1 && col < cols; col++ {
			adj = append(adj, position{row, col})
		}
	}
	return adj
}

func inPath(path []position, pos position) bool {
	for _, p := range path {
		if pos == p {
			return true
		}
	}
	return false
}
