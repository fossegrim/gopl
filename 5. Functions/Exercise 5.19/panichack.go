// panichack doubles a number using a novel panic/defer/recover hack.
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: panichack int")
		os.Exit(1)
	}
	x, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer func() {
		// This doesn't handle the scenario where another, or no, panic occurs in double.
		// I don't think that scenario is very likely though...

		fmt.Println(recover())
	}()
	double(x)
}

// double panics with the double of x to "return" it to a caller using defer/recover to catch it.
// Note that this is obviously not the proper way to return values in Go :-).
func double(x int) {
	panic(x * 2)
}
