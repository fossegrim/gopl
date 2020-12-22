package main

import (
	"fmt"

	"./rotate"
)

// Demonstrate that the rotate.Rotate function works correctly.
func main() {
	a := []int{1,2,3,4,5,6,7,8,9,10}
	for rotations := 0; rotations <= 100; rotations++ {
		arotated := make([]int, len(a))
		copy(arotated, a)
		rotate.Rotate(arotated, rotations)
		fmt.Printf("%v rotated by %3d rotations is %2v.\n", a, rotations, arotated)
	}

	a = []int{1,2,3}
	for rotations := 0; rotations <= 30; rotations++ {
		arotated := make([]int, len(a))
		copy(arotated, a)
		rotate.Rotate(arotated, rotations)
		fmt.Printf("%v rotated by %d rotations is %v.\n", a, rotations, arotated)
	}
}
