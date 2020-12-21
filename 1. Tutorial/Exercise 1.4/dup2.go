// Dup2 prints the count and text of lines that appear more than once
// in the input. It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//                    filename   line    includes
	fileLines := make(map[string]map[string]bool)
	//                 line   count
	counts := make(map[string]int)
	filenames := os.Args[1:]
	if len(filenames) == 0 {
		// "" is used to represent stdin. A conflict with a on-disk filename is impossible since they must be at least one character long.
		filename := ""
		countLines(filename, os.Stdin, counts, fileLines)
	} else {

		for _, filename := range filenames {
			f, err := os.Open(filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			}
			countLines(filename, f, counts, fileLines)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			// only explicitly specify which file(s) contained the lines if it's necesarry
			if len(filenames) == 0 || len(filenames) == 1 {
				fmt.Printf("%d\t%s\n", n, line)
			} else {
				filesContainingLine, sep := "", ""
				for filename, lines := range fileLines {
					if lines[line] {
						filesContainingLine += sep + filename
						sep = ","
					}
				}
				fmt.Printf("%s\t%d\t%s\n", filesContainingLine, n, line)
			}
		}
	}
}

func countLines(filename string, f *os.File, counts map[string]int, fileLines map[string]map[string]bool) {
	input := bufio.NewScanner(f)
	fileLines[filename] = make(map[string]bool)
	for input.Scan() {
		counts[input.Text()]++
		fileLines[filename][input.Text()] = true
	}
	// NOTE: ignoring potential erros from input.Err()
}
