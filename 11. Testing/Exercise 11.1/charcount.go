package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"./countcharacters"
)

func main() {
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
		os.Exit(1)
	}
	characterCounts, utflen, invalid, categoryCounts := countcharacters.CountCharacters(string(bytes))

	fmt.Printf("category\tcount\n")
	for c, n := range categoryCounts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range characterCounts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
