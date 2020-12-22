package main

import (
	"fmt"

	"./reverse"
)

// Demonstrate that reverse.Reverse works correctly.
func main() {
	a := [10]int{2}
	for i := 1; i < len(a); i++ {
		a[i] = a[i-1] *2
	}
	
	areverse := a
	reverse.Reverse(&areverse)
	fmt.Printf("The reverse of %v is %v.\n", a, areverse)

	a = areverse
	reverse.Reverse(&areverse)
	fmt.Printf("The reverse of %v is %v.\n", a, areverse)
}
