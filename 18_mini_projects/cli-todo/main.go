// Command cli-todo is a tiny JSON-backed todo list. It demonstrates the
// load -> mutate -> save cycle over a single JSON file, including the atomic
// "write temp + rename" pattern so a crash mid-write can't corrupt the list.
//
// Usage:
//
//	go run ./18_mini_projects/cli-todo add "learn go"
//	go run ./18_mini_projects/cli-todo list
//	go run ./18_mini_projects/cli-todo done 1
//
// The list is stored in todo.json in the current directory (override with
// -file). Each run reads the whole file, applies one change, and writes it back.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

// Todo is a single list item. The struct tags control the JSON field names so
// the on-disk format stays stable even if we rename the Go fields later.
type Todo struct {
	Text string `json:"text"`
	Done bool   `json:"done"`
}

// load reads and decodes the todo file. A missing file is not an error: it just
// means an empty list, which makes the very first run work with no setup.
func load(path string) ([]Todo, error) {
	b, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return nil, nil // first run — start with an empty list
	}
	if err != nil {
		return nil, err
	}
	var xs []Todo
	// Unmarshal returns an error for malformed JSON; we surface it to the caller.
	return xs, json.Unmarshal(b, &xs)
}

// save writes the list back to disk atomically. We never write `path` directly:
// a crash partway through would leave a half-written, corrupt file. Instead we
// write a temp file in the SAME directory, then os.Rename it over the target.
// Rename within one filesystem is atomic, so readers always see either the old
// file or the complete new one — never a partial write.
func save(path string, xs []Todo) error {
	b, err := json.MarshalIndent(xs, "", "  ")
	if err != nil {
		return err
	}

	// Temp file must share the target's directory so the final Rename stays on
	// the same filesystem (cross-device rename is not atomic and may fail).
	dir := filepath.Dir(path)
	tmp, err := os.CreateTemp(dir, ".todo-*.tmp")
	if err != nil {
		return err
	}
	tmpName := tmp.Name()
	// If anything below fails, make sure we don't leave the temp file behind.
	defer os.Remove(tmpName)

	if _, err := tmp.Write(b); err != nil {
		tmp.Close()
		return err
	}
	// Close before Rename so all data is flushed to the OS and the handle is
	// released (important on Windows, harmless on Unix).
	if err := tmp.Close(); err != nil {
		return err
	}

	// The atomic swap. After this returns, `path` is the new content.
	return os.Rename(tmpName, path)
}

func main() {
	// -file lets tests / users point at a different list; default todo.json.
	file := flag.String("file", "todo.json", "path to the todo JSON file")
	flag.Parse()

	// flag.Args() returns only the positional args left AFTER flag parsing — it
	// excludes the program name and any parsed flags (unlike os.Args, whose [0]
	// is the binary). So args[0] is the subcommand (add/list/done) and the rest
	// are its operands.
	args := flag.Args()
	if len(args) == 0 {
		usage()
	}

	xs, err := load(*file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading %s: %v\n", *file, err)
		os.Exit(1)
	}

	switch args[0] {
	case "add":
		// Everything after "add" becomes the todo text.
		if len(args) < 2 {
			fmt.Fprintln(os.Stderr, "add: missing todo text")
			os.Exit(1)
		}
		text := args[1]
		xs = append(xs, Todo{Text: text})
		mustSave(*file, xs)
		fmt.Printf("added #%d: %s\n", len(xs), text)

	case "list":
		// Print the list with 1-based indices and a [ ] / [x] status box.
		if len(xs) == 0 {
			fmt.Println("(no todos)")
			return
		}
		for i, t := range xs {
			box := " "
			if t.Done {
				box = "x"
			}
			fmt.Printf("%d. [%s] %s\n", i+1, box, t.Text)
		}

	case "done":
		// Mark the item at the given 1-based index as done.
		if len(args) < 2 {
			fmt.Fprintln(os.Stderr, "done: missing item number")
			os.Exit(1)
		}
		n, err := strconv.Atoi(args[1])
		if err != nil || n < 1 || n > len(xs) {
			fmt.Fprintf(os.Stderr, "done: invalid item number %q\n", args[1])
			os.Exit(1)
		}
		xs[n-1].Done = true
		mustSave(*file, xs)
		fmt.Printf("done #%d: %s\n", n, xs[n-1].Text)

	default:
		usage()
	}
}

// mustSave saves or exits with a clear error — the mutating commands can't
// meaningfully continue if persistence fails.
func mustSave(path string, xs []Todo) {
	if err := save(path, xs); err != nil {
		fmt.Fprintf(os.Stderr, "error saving %s: %v\n", path, err)
		os.Exit(1)
	}
}

func usage() {
	fmt.Fprintln(os.Stderr, "usage: cli-todo [-file path] <add <text> | list | done <n>>")
	os.Exit(2)
}
