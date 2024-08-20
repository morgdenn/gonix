package main

import (
	_ "embed"
	"flag"
	"fmt"
	"os"
	"text/tabwriter"
	"time"
)

//go:embed help.txt
var helpText string

func main() {

	// Define the --help flag
	flag.Usage = func() {
		fmt.Print(helpText)
	}

	// Define the -l flag
	longFormat := flag.Bool("l", false, "use a long listing format")

	// Parse command-line flags
	flag.Parse()

	// Get the directory to list; default to current directory if not provided
	dir := "."
	if flag.NArg() > 0 {
		dir = flag.Arg(0)
	}

	// Read the directory
	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Error reading the directory", err)
	}

	if *longFormat {
		fmt.Println("Total ?")

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
		for _, entry := range entries {

			// Get the info for the entry
			info, err := entry.Info()
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}

			size := info.Size()
			month := info.ModTime().Format("Jan")
			day := info.ModTime().Day()
			dateEnd := getDateEnd(info)
			name := entry.Name()

			fmt.Fprintf(w, "%s\t", info.Mode())
			fmt.Fprintf(w, "%d\t", size)
			fmt.Fprintf(w, "%s\t", month)
			fmt.Fprintf(w, "%2d\t", day)
			fmt.Fprintf(w, "%s\t", dateEnd)
			fmt.Fprintf(w, "%s\t", name)
			fmt.Fprintf(w, "\n")
		}
		w.Flush()

		return
	}

	// Default behavior: print file names
	for _, entry := range entries {
		fmt.Print(entry.Name() + "  ")
	}
	fmt.Println()
}

func getDateEnd(info os.FileInfo) string {
	// Calculate the time six months ago
	sixMonthsAgo := time.Now().AddDate(0, -6, 0)

	// Check if the file is less than six months old
	if info.ModTime().After(sixMonthsAgo) {
		return info.ModTime().Format("15:04")
	}
	return info.ModTime().Format(" 2006")
}
