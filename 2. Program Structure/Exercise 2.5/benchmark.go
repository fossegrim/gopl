// Benchmark the performace of popcount.{Flat,Looped}PopCount
package main

import (
	"fmt"
	"time"

	"./popcount"
)

const ncalls = 1536 // cherry picked number ;)

func main() {
	start := time.Now()
	for i := uint64(0); i < ncalls; i++ {
		popcount.FlatPopCount(i)
	}
	fmt.Printf("%d calls to FlatPopCount took %v\n", ncalls, time.Since(start))

	start = time.Now()
	for i := uint64(0); i < ncalls; i++ {
		popcount.LoopedPopCount(i)
	}
	fmt.Printf("%d calls to LoopedPopCount took %v\n", ncalls, time.Since(start))

	start = time.Now()
	for i := uint64(0); i < ncalls; i++ {
		popcount.NoTablePopCount(i)
	}
	fmt.Printf("%d calls to NoTablePopCount took %v\n", ncalls, time.Since(start))

	start = time.Now()
	for i := uint64(0); i < ncalls; i++ {
		popcount.EpicPopCount(i)
	}
	fmt.Printf("%d calls to EpicPopCount took %v\n", ncalls, time.Since(start))
}
