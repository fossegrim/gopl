// Demonstrate that the anagram.IsAnagram function behaves correctly
package main

import (
	"fmt"

	"./anagram"
)

func main() {
	// Simple true case
	demonstrateIsAnagram("123", "321")
	// Extra character in one of the strings
	demonstrateIsAnagram("1234", "321")
	demonstrateIsAnagram("123", "4321")
	// Too many of one character
	demonstrateIsAnagram("1123", "321")
	demonstrateIsAnagram("123", "3211")
	// Equal bytes, but unequal runes
	demonstrateIsAnagram("\xe4\xb8\x96\xe7\x95\x8c", "\xe4\x95\x8c\xe7\xb8\x96")
}

// demonstrateIsAnagram demonstrates what anagram.IsAnagram returns with a pair of arguments
func demonstrateIsAnagram(s1, s2 string) {
	fmt.Printf("IsAnagram(%q, %q) = %t\n", s1, s2, anagram.IsAnagram(s1, s2))
}
