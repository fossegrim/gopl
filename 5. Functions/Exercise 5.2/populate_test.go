package populate

import (
	"fmt"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestPopulateCountMap(t *testing.T) {
	wantedCountMap := map[string]int{
		"html":  1,
		"h1":    1,
		"p":     2,
		"style": 1,
	}

	s := `
<!doctype html>
<html>
  <head>
    <title></title>
    <meta charset="utf-8" />
    <link rel="stylesheet" href="style.css"/>
  </head>
  <body>
    <h1>Olav's Website</h1>
    <p>Welcome to my epic site
    <p>There is nothing here
    <style>
      body {
	  padding: 0;
      }
    </style>
  </body>
</html>`
	n, _ := html.Parse(strings.NewReader(s))
	countMap := map[string]int{
		"html":  0,
		"h1":    0,
		"p":     0,
		"style": 0,
	}
	countMap = PopulateCountMap(n, countMap)

	if fmt.Sprint(countMap) != fmt.Sprint(wantedCountMap) {
		t.Errorf("PopulateCountMap returns %v, want %v", countMap, wantedCountMap)
	}
}
