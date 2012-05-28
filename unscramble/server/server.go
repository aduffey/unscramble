package main

import (
	"appengine"
	"bufio"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"runtime/debug"
	"unscramble/solver"
)

const (
	templateFile = "templates/index.html"
	dictFile     = "word.lst"
	prefix       = "/"
)

// Globals. These will get initialized in main() and will not change afterwards.
var (
	dict *solver.Dict
	tmpl *template.Template
)

func serveTemplate(w http.ResponseWriter, tmpl *template.Template,
	ctx interface{}) {
	w.Header().Set("Content-type", "text/html")
	tmpl.Execute(w, ctx)
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Handle panics so we can at least return a response code
	defer func() {
		if err := recover(); err != nil {
			appCtx := appengine.NewContext(r)
			appCtx.Errorf("%s: %s", err, debug.Stack())
			http.Error(w, "Internal server error",
				http.StatusInternalServerError)
		}
	}()

	path := r.URL.Path[len(prefix):]
	if path == "" {
		// Go to the main page
		ctx := emptyContext()
		serveTemplate(w, tmpl, ctx)
	} else if board, err := solver.NewBoardFromString(path); err == nil {
		// Path represents a valid board string, show the solutions
		solutions := solver.Solve(board, dict)
		ctx := newContext(board, solutions)
		serveTemplate(w, tmpl, ctx)
	} else {
		http.NotFound(w, r)
	}
}

// loadDict loads the dictionary or panics on errors.
func loadDict(filename string) *solver.Dict {
	f, fileErr := os.Open(filename)
	if fileErr != nil {
		panic(fmt.Sprintf("Error opening dictionary file: %s", fileErr))
	}
	defer f.Close()

	dict := solver.NewDict()
	reader := bufio.NewReader(f)
	for lineNbr := 0; ; lineNbr++ {
		line, isPrefix, readErr := reader.ReadLine()
		if isPrefix {
			panic(fmt.Sprintf("Dictionary line %d is too long", lineNbr))
		} else if readErr != nil && readErr != io.EOF {
			panic(fmt.Sprintf("Error reading dictionary: %s", readErr))
		}
		if line == nil {
			break
		}
		word := string(line)
		if !solver.ValidString(word) {
			panic(fmt.Sprintf("Dictionary word \"%s\" at line %d is invalid",
				word, lineNbr))
		}
		if !dict.Add(word) {
			panic(fmt.Sprintf(
				"Dictionary contains duplicate word \"%s\" at line %d", word,
				lineNbr))
		}
	}

	return dict
}

func init() {
	log.Print("Initializing...\n")

	dict = loadDict(dictFile)

	var err error
	tmpl, err = template.ParseFiles(templateFile)
	if err != nil {
		panic(fmt.Sprintf("Could not load template from file %s: %s\n",
			templateFile, err))
	}

	http.HandleFunc(prefix, handler)
}
