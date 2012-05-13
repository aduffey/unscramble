package main

import (
	"solver"
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

func loadDict(filename string) (*solver.Dict, error) {
	f, fileErr := os.Open(filename)
	if fileErr != nil {
		return nil, fileErr
	}
	defer f.Close()

	dict := solver.NewDict()
	reader := bufio.NewReader(f)
	for {
		line, isPrefix, readErr := reader.ReadLine()
		if isPrefix || (readErr != nil && readErr != io.EOF) {
			return nil, readErr
		}
		if line == nil {
			break
		}
		// TODO(aduffey) check if valid
		dict.Add(string(line))
	}

	return dict, nil
}

type parseError string

func (pe parseError) Error() string {
	return string(pe)
}

func parseBoard(boardString string) (*solver.Board, error) {
	b := &solver.Board{}
	row, col := 0, 0
	for index, char := range boardString {
		if row >= solver.Rows || col >= solver.Cols {
			err := fmt.Sprintf("Input string is too long, expected %d cells",
				solver.Rows*solver.Cols)
			return nil, parseError(err)
		}

		if char == '2' {
			if b.Modifiers[row][col] == solver.None {
				b.Modifiers[row][col] = solver.X2Letter
			} else if b.Modifiers[row][col] == solver.X2Letter {
				b.Modifiers[row][col] = solver.X2Word
			} else {
				err := fmt.Sprintf("Unexpected \"2\" at index %v", index)
				return nil, parseError(err)
			}
		} else if char == '3' {
			if b.Modifiers[row][col] == solver.None {
				b.Modifiers[row][col] = solver.X3Letter
			} else if b.Modifiers[row][col] == solver.X3Letter {
				b.Modifiers[row][col] = solver.X3Word
			} else {
				err := fmt.Sprintf("Unexpected \"3\" at index %d", index)
				return nil, parseError(err)
			}
		} else if solver.ValidChar(char) {
			b.Chars[row][col] = char
			col++
			if col >= solver.Cols {
				row++
				col = 0
			}
		} else {
			err := fmt.Sprintf("Invalid symbol %v at index %d", char, index)
			return nil, parseError(err)
		}
	}
	if row < (solver.Rows-1) || row == (solver.Rows-1) &&
		col < (solver.Cols-1) {
		err := fmt.Sprintf("Input string is too short, expected %d cells",
			solver.Rows*solver.Cols)
		return nil, parseError(err)
	}
	return b, nil
}

func main() {
	if len(os.Args) != 3 {
		fmt.Print("Error: invalid arguments\n")
		os.Exit(1)
	}
	dictFile := os.Args[1]
	boardString := os.Args[2]

	fmt.Print("Loading dictionary...\n")
	loadStart := time.Now()
	dict, dictErr := loadDict(dictFile)
	if dictErr != nil {
		fmt.Printf("Error loading dictionary: %s\n", dictErr.Error())
		os.Exit(1)
	}
	loadElapsed := float32(time.Since(loadStart)) / float32(time.Millisecond)
	fmt.Printf("Loaded dictionary in %.2fms\n", loadElapsed)

	board, boardErr := parseBoard(boardString)
	if boardErr != nil {
		fmt.Printf("Error parsing board: %s\n", boardErr.Error())
		os.Exit(1)
	}

	fmt.Print("Solving...\n")
	solveStart := time.Now()
	sols := solver.Solve(board, dict)
	solveElapsed := float32(time.Since(solveStart)) / float32(time.Millisecond)
	fmt.Printf("Solved board in %.2fms\n", solveElapsed)
	for _, sol := range sols {
		fmt.Printf("%v\n", sol)
	}
}
