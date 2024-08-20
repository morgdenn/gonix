package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: which <command>")
		os.Exit(1)
	}

	command := os.Args[1]

	// Get the PATH environment variable
	pathEnv := os.Getenv("PATH")

	// Split the PATH by the operating system's path separator
	paths := strings.Split(pathEnv, string(os.PathListSeparator))

	// Get the PATHEXT environment variable
	// PATHEXT contains extensions for executable files
	exts := getExecutableExtensions()

	// Iterate through each directory in the PATH
	for _, dir := range paths {
		for _, ext := range exts {
			// Construct the full path to the command
			fullPath := filepath.Join(dir, command+ext)

			// Check if the command exists and is executable
			if fileExists(fullPath) {

				correctCasePath := findCorrectCasePath(dir, command, ext)

				fmt.Println(correctCasePath)
				return
			}
		}
	}

	os.Exit(1)
}

// fileExists checks if a file exists and is executable
func fileExists(path string) bool {

	info, err := os.Stat(path)

	if err != nil {
		return false
	}

	if info.IsDir() {
		return false
	}

	return true
}

// getExecutableExtensions returns the list of executable extensions
func getExecutableExtensions() []string {
	// Get the PATHEXT environment variable, which lists executable file extensions
	pathext := os.Getenv("PATHEXT")

	if pathext == "" {
		// Default to common executable extensions if PATHEXT is not set
		return []string{""}
	}

	// Split PATHEXT by the semicolon separator and add an empty string for commands without extension
	exts := strings.Split(pathext, ";")
	exts = append([]string{""}, exts...)

	return exts
}

// findCorrectCasePath searches for the command in a directory and returns the correct path
func findCorrectCasePath(dir, command string, ext string) string {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return ""
	}

	for _, entry := range entries {

		if entry.IsDir() {
			continue
		}

		// Compare file names case-insensitively
		if strings.EqualFold(entry.Name(), command+ext) {
			// Construct the full path with the correct case
			return filepath.Join(dir, entry.Name())
		}
	}

	return filepath.Join(dir, command, ext)
}
