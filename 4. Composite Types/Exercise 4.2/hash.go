package main

import (
	"crypto/sha512"
	"fmt"
	"io/ioutil"
	"os"
)

// Print the SHA{256,384,512) hash of standard input.
func main() {
	var hashingAlgorithm string
	switch len(os.Args) {
	case 1:
		hashingAlgorithm = "SHA-256"
	case 2:
		hashingAlgorithm = os.Args[1]
	default:
		exitWithUsage()
	}

	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	// This switch switch block is very WET.
	// In each case, the hash variable is a byte array of different lengths.
	// We cannot declare a hash variable in the outer scope, set it in the switch cases and print it afterwards, because the return types of Sum{512_256,384,512} are different.
	// I think a better way of doing this will be covered later in the book, possibly in the 7. Interfaces chapter.
	switch hashingAlgorithm {
	case "SHA-256":
		hash := sha512.Sum512_256(input)
		fmt.Printf("%x\n", hash)
	case "SHA-384":
		hash := sha512.Sum384(input)
		fmt.Printf("%x\n", hash)
	case "SHA-512":
		hash := sha512.Sum512(input)
		fmt.Printf("%x\n", hash)
	default:
		exitWithUsage()
	}
}

func exitWithUsage() {
	fmt.Fprintln(os.Stderr, "USAGE: hash [SHA-256|SHA-384|SHA-512]")
	os.Exit(1)
}
