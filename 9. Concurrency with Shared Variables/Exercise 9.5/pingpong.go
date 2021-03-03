package main

import (
	"fmt"
	"time"
)

func main() {
	pings, pongs := make(chan struct{}), make(chan struct{})
	nPings, nPongs := 0, 0

	exit := false
	go func() {
		for {
			if exit {
				return
			}
			select {
			case pings <- struct{}{}:
				nPings++
			case <-pongs:
				// Do nothing.
			}
		}
	}()
	go func() {
		for {
			if exit {
				return
			}
			select {
			case pongs <- struct{}{}:
				nPongs++
			case <-pings:
				// Do nothing.
			}
		}
	}()
	seconds := 10
	time.Sleep(time.Duration(1000 * 1000 * 1000 * seconds))
	exit = true
	close(pings)
	close(pongs)
	fmt.Printf("%d pings and %d pongs\n", nPings, nPongs)
	fmt.Printf("%f pings per second\n", float64(nPings)/float64(seconds))
	fmt.Printf("%f pongs per second\n", float64(nPongs)/float64(seconds))
}
