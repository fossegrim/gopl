package count

import (
	"bufio"
	"strings"
	
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// countWordsAndImages counts the number of words and images in the html node n.
func countWordsAndImages(n *html.Node) (words, images int) {
	if n.Type == html.TextNode {
		s := bufio.NewScanner(strings.NewReader(n.Data))
		s.Split(bufio.ScanWords)
		for s.Scan() {
			words++	
		}
	} else if n.DataAtom == atom.Img {
		images++
	} 

	for child := n.FirstChild; child != nil; child = child.NextSibling {
		w, i := countWordsAndImages(child)
		words += w
		images += i
	}

	return
}
