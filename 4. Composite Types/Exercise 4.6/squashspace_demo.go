package main

import (
	"fmt"			 

	"./squashspace"
)
// Demonstrate that the squashspace.Squash function behaves correctly
func main() {
	s :=  []byte(`
this is a string with newlines


		tabs
	
   and    lot of spaces






when this is desquashed all    sequences of white space will be   squashed to single spaces








s


heres some more text



`)
	sSquashed := make([]byte, len(s))
	copy(sSquashed, s)
	squashspace.Squash(&sSquashed)

	fmt.Printf("%q squashed is %q.\n", s, sSquashed)
}
