package main

import (
	"bufio"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"unscramble/solver"
)

const (
	templateFile  = "templates/index.html"
	dictFile      = "word.lst"
	prefix        = "/"
	staticRoot    = "static/"
	listenAddress = ":8080"
)

// Globals. These will get initialized in main() and will not change afterwards.
var (
	dict *solver.Dict
	tmpl *template.Template
)

func fileExists(name string) bool {
	_, err := os.Stat(name)
	return err == nil || os.IsExist(err)
}

func serveTemplate(w http.ResponseWriter, tmpl *template.Template,
	ctx interface{}) {
	w.Header().Set("Content-type", "text/html")
	tmpl.Execute(w, ctx)
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Handle panics so we can at least return a response code
	defer func() {
		if err := recover(); err != nil {
			log.Printf("ERROR: %s\n", err)
			log.Printf("    At url: %s", r.URL)
			http.Error(w, "Internal server error",
				http.StatusInternalServerError)
		}
	}()

	path := r.URL.Path[len(prefix):]
	if path == "" {
		// Go to the main page
		ctx := emptyContext()
		serveTemplate(w, tmpl, ctx)
	} else if file := staticRoot + path; fileExists(file) {
		// If the path matches a static file, serve it up
		http.ServeFile(w, r, file)
	} else if board, err := solver.NewBoardFromString(path); err == nil {
		// Path represents a valid board string, show the solutions
		solutions := solver.Solve(board, dict)
		ctx := newContext(board, solutions)
		serveTemplate(w, tmpl, ctx)
	} else {
		http.NotFound(w, r)
	}
}

// TODO(aduffey) maybe move this into solver?
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
	log.Print("Initializing...\n")

	var err error
	// Load the dictionary
	dict, err = loadDict(dictFile)
	if err != nil {
		log.Fatalf("Could not load dictionary from file %s: %s\n", dictFile,
			err)
	}
	log.Printf("Successfully loaded dictionary from file %s\n", dictFile)

	// Load the template
	tmpl, err = template.ParseFiles(templateFile)
	if err != nil {
		log.Fatalf("Could not load template from file %s: %s\n", templateFile,
			err)
	}
	log.Printf("Successfully loaded template from file %s\n", templateFile)

	// Serve
	http.HandleFunc(prefix, handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
