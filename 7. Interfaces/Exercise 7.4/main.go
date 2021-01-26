package main

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html"
)

var htmlString = `<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>demo</title>
  </head>
  <body>
    <p> hei
  </body>
</html>`

type Reader struct {
	s string
}

func (r *Reader) Read(p []byte) (n int, err error) {
	n = copy(p, r.s)
	r.s = r.s[n:]
	if n == len(r.s) {
		err = io.EOF
	}
	return
}

func NewReader(s string) *Reader {
	return &Reader{s}
}

func main() {
	n, _ := html.Parse(NewReader(htmlString))
	html.Render(os.Stdout, n)
	fmt.Println()
}
