// Benchmark the performace of popcount.{Flat,Looped}PopCount
package main

import (
	"fmt"
	"time"

	"./popcount"
)

const ncalls = 100000000

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
}
