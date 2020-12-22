package main

import (
	"fmt"

	"./dedup"
)

// Demonstrate that dedup.Dedup function works correctly.
func main() {
	a := []string{"olav", "olav", "olav", "fosse", "fossegrim", "fossegrim"}
	adeduped := make([]string, len(a))
	copy(adeduped, a)
	dedup.Dedup(adeduped)
	fmt.Printf("%v deduped is %v.\n", a, adeduped)
}
