// report frequency of each word
// input.Split(bufio.ScanWords) to break input into words

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords) // scan word word by word
	
	wordsCount := make(map[string]int)
	for input.Scan() {
		s := input.Text()
		wordsCount[s]++
	}
	err := input.Err()
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	println("word\tcount")
	for word, count := range wordsCount {
		fmt.Printf("%s\t%d\n", word, count)
	}
}
