package main

import (
	"os"
	"os/exec"
)

func main() {
	// Base command: "ls -l"
	args := append([]string{"-l"}, os.Args[1:]...)

	// Create the command with all passed arguments
	cmd := exec.Command("ls", args...)

	// Set the standard input, output, and error to the same as the parent process
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run the command
	err := cmd.Run()
	if err != nil {
		// Print any errors that occur
		os.Exit(1)
	}
}
