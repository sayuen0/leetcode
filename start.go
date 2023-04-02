package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "error: Empty Problem name")
		os.Exit(1)
	}

	input := strings.Join(os.Args[1:], " ")
	splitInput := strings.SplitN(input, ".", 2)

	if len(splitInput) != 2 {
		fmt.Fprintln(os.Stderr, "error: Invalid input format. Please use the correct problem name format.")
		os.Exit(1)
	}

	first, err := strconv.Atoi(strings.TrimSpace(splitInput[0]))
	if err != nil {
		fmt.Fprintln(os.Stderr, "error: Invalid number in the problem name.")
		os.Exit(1)
	}

	second := strings.TrimSpace(splitInput[1])
	snakeCaseSecond := toSnakeCase(second)

	d := fmt.Sprintf("no_%03d", first)
	f := snakeCaseSecond

	createFiles(d, f)
	createFiles(filepath.Join(d, "example"), f)
}

func toSnakeCase(s string) string {
	words := strings.Fields(s)
	for i, word := range words {
		words[i] = strings.ToLower(string(word[0])) + string(word[1:])
	}
	return strings.Join(words, "_")
}

func createFiles(dir string, fileName string) {
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error: Unable to create directory")
		os.Exit(1)
	}

	err = os.WriteFile(filepath.Join(dir, fileName+".go"), []byte("package main\n"), 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error: Unable to create .go file")
		os.Exit(1)
	}

	err = os.WriteFile(filepath.Join(dir, fileName+"_test.go"), []byte("package main\n"), 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error: Unable to create _test.go file")
		os.Exit(1)
	}
}
