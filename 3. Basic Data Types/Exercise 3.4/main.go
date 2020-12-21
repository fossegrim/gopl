// Serve SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"log"
	"net/http"

	"./surface"
)

const addr = "localhost:8000"

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")
		surface.Render(w)
	}
	http.HandleFunc("/", handler)
	fmt.Printf("listening on: %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
