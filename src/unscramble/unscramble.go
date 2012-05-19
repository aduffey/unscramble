package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"solver"
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

	board, boardErr := solver.NewBoardFromString(boardString)
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
