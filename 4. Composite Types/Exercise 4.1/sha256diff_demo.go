package main

import (
	"crypto/sha256"
	"fmt"

	"./sha256diff"
)

// Demonstrate sha256diff.Count
func main() {
	left := "left"
	right := "right"

	leftHash := sha256.Sum256([]byte(left))
	rightHash := sha256.Sum256([]byte(right))

	differingBytes := sha256diff.Count(leftHash, rightHash)
	fmt.Printf("%q's SHA256 hash is %x\n", left, leftHash)
	fmt.Printf("%q's SHA256 hash is %x\n", right, rightHash)
	fmt.Printf("%q's and %q's SHA256 hashes have %d differing bits\n", left, right, differingBytes)
}
