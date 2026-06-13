package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type category struct {
	Folder string `json:"folder"`
	Title  string `json:"title"`
	Count  int    `json:"count"`
	Tasks  []task `json:"-"`
}

type task struct {
	Category string `json:"category"`
	Title    string `json:"title"`
	File     string `json:"file"`
}

type manifest struct {
	TotalExamples int        `json:"total_examples"`
	Categories    []category `json:"categories"`
	Examples      []task     `json:"examples"`
}

func main() {
	categories, err := discoverCategories(".")
	if err != nil {
		fatal(err)
	}

	var allTasks []task
	for _, category := range categories {
		allTasks = append(allTasks, category.Tasks...)
	}

	if err := writeREADME(categories, len(allTasks)); err != nil {
		fatal(err)
	}
	if err := writeCatalog(categories, len(allTasks)); err != nil {
		fatal(err)
	}
	if err := writeManifest(categories, allTasks); err != nil {
		fatal(err)
	}
}

func discoverCategories(root string) ([]category, error) {
	entries, err := os.ReadDir(root)
	if err != nil {
		return nil, err
	}
	var categories []category
	for _, entry := range entries {
		if !entry.IsDir() || len(entry.Name()) < 3 || entry.Name()[2] != '_' {
			continue
		}
		folder := entry.Name()
		title, err := firstHeading(filepath.Join(root, folder, "README.md"))
		if err != nil {
			return nil, fmt.Errorf("read %s title: %w", folder, err)
		}
		files, err := filepath.Glob(filepath.Join(root, folder, "*.md"))
		if err != nil {
			return nil, err
		}
		sort.Strings(files)
		var tasks []task
		for _, file := range files {
			if filepath.Base(file) == "README.md" {
				continue
			}
			name := strings.TrimSuffix(filepath.Base(file), ".md")
			tasks = append(tasks, task{Category: folder, Title: name, File: filepath.ToSlash(file)})
		}
		categories = append(categories, category{Folder: folder, Title: title, Count: len(tasks), Tasks: tasks})
	}
	sort.Slice(categories, func(i, j int) bool { return categories[i].Folder < categories[j].Folder })
	return categories, nil
}

func firstHeading(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	for _, line := range strings.Split(string(data), "\n") {
		if strings.HasPrefix(line, "# ") {
			return strings.TrimSpace(strings.TrimPrefix(line, "# ")), nil
		}
	}
	return "", fmt.Errorf("missing H1 heading")
}

func writeREADME(categories []category, total int) error {
	var b strings.Builder
	fmt.Fprintf(&b, "# Golang live interview tasks\n\n")
	fmt.Fprintf(&b, "A collection of **%d** live-coding and discussion tasks covering the Go language, standard library, concurrency, HTTP, scraping, data access, testing, performance, and service design.\n\n", total)
	b.WriteString("Each task contains a prompt, concepts or requirements, a candidate solution or solution outline, and interview pitfalls. Most tasks use only the standard library; Colly tasks require `github.com/gocolly/colly/v2`.\n\n")
	b.WriteString("- [Complete task catalog](./TASK_CATALOG.md)\n")
	b.WriteString("- [Official learning and API references](./REFERENCES.md)\n")
	b.WriteString("- [Machine-readable manifest](./manifest.json)\n\n")
	b.WriteString("## Categories\n\n")
	for _, category := range categories {
		fmt.Fprintf(&b, "- [%s](./%s/README.md) - %s (%d tasks)\n", category.Folder, category.Folder, category.Title, category.Count)
	}
	b.WriteString("\n## Regenerate indexes\n\n```bash\ngo run ./tools/generate_catalog.go\n```\n")
	return os.WriteFile("README.md", []byte(b.String()), 0o644)
}

func writeCatalog(categories []category, total int) error {
	var b strings.Builder
	b.WriteString("# Task catalog\n\n")
	fmt.Fprintf(&b, "Generated list of %d live interview tasks. Run `go run ./tools/generate_catalog.go` after adding or removing task files.\n\n", total)
	for _, category := range categories {
		fmt.Fprintf(&b, "## %s - %s\n\n", category.Folder, category.Title)
		for _, task := range category.Tasks {
			fmt.Fprintf(&b, "- [%s](./%s)\n", task.Title, task.File)
		}
		b.WriteByte('\n')
	}
	return os.WriteFile("TASK_CATALOG.md", []byte(b.String()), 0o644)
}

func writeManifest(categories []category, tasks []task) error {
	data, err := json.MarshalIndent(manifest{TotalExamples: len(tasks), Categories: categories, Examples: tasks}, "", "  ")
	if err != nil {
		return err
	}
	data = append(data, '\n')
	return os.WriteFile("manifest.json", data, 0o644)
}

func fatal(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
