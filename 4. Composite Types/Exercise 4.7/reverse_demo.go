package main

import (
	"fmt"

	"./reverse"
)

// Demonstrate that reverse.Reverse works correctly.
func main() {
	// ASCII
	s := []byte("lorem")
	sreverse := make([]byte, len(s))
	copy(sreverse, s)
	reverse.Reverse(sreverse)
	fmt.Printf("The reverse of %q is %q.\n", s, sreverse)

	// Uneven number of UTF-8 characters
	s = []byte("æøå")
	sreverse = make([]byte, len(s))
	copy(sreverse, s)
	reverse.Reverse(sreverse)
	fmt.Printf("The reverse of %q is %q.\n", s, sreverse)

	// Even number of UTF-8 characters
	s = []byte("æøåæøå")
	sreverse = make([]byte, len(s))
	copy(sreverse, s)
	reverse.Reverse(sreverse)
	fmt.Printf("The reverse of %q is %q.\n", s, sreverse)

	// Many different ASCII and non-ASCII UTF-8 characters
	s = []byte("Thøre are mæny cool characters. Søm in norweigan and other in other language (hello, 世界)")
	sreverse = make([]byte, len(s))
	copy(sreverse, s)
	reverse.Reverse(sreverse)
	fmt.Printf("The reverse of %q is %q.\n", s, sreverse)
}
