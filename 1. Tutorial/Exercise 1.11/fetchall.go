// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "USAGE:\tfetchall file URL ...")
		os.Exit(1)
	}
	filename := os.Args[1]
	f, err := os.Create(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetchall: %v\n", err)
		os.Exit(1)
	}
	urls := os.Args[2:]
	for _, url := range urls {
		go fetch(f, url, ch) // start a goroutine
	}
	for range urls {
		fmt.Println(<-ch) // receive from channel ch
	}
	f.Close()
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(f *os.File, url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	nbytes, err := io.Copy(f, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		fmt.Sprintf("while reading %s: %v", url, err)
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
