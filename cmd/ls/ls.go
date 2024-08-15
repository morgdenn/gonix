package main

import (
	"fmt"
	"os"
	"strings"
)

var longestName = 0

func main() {
	fmt.Println("Files...")

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting the current Dir", err)
	}

	fmt.Println(" Directory of: ", dir)

	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Error reading the directory", err)
	}

	directories := make([]os.DirEntry, 0)
	files := make([]os.DirEntry, 0)

	for _, entry := range entries {
		if len(entry.Name()) > longestName {
			longestName = len(entry.Name())
		}

		if entry.IsDir() {
			directories = append(directories, entry)
		} else {
			files = append(files, entry)
		}
	}

	for _, entry := range directories {
		outputDir(entry)
	}

	for _, entry := range files {
		outputFile(entry)
	}

}

func outputDir(entry os.DirEntry) {
	padding := getPadding(entry.Name())
	fmt.Println("/ ", entry.Name(), padding, "| -")
}

func outputFile(entry os.DirEntry) {
	padding := getPadding(entry.Name())
	fmt.Println("- ", entry.Name(), padding, "| -")
}

func getPadding(name string) string {
	padding := longestName - len(name)
	return strings.Repeat(" ", padding)
}
