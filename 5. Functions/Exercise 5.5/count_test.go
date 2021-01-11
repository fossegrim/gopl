package count

import (
	"testing"
	"strings"

	"golang.org/x/net/html"
)

// The test should use a html tree parsed with html.Parse(r io.Reader) (n *html.Node, error)
func TestCountWordsAndImages(t *testing.T) {
	wantedWords := 26
	wantedImages := 2

	s := `
<html>
	<head>
		<title>testing test</title>
	</head>
	<body>
		<h1>This is a nice header</h1>
		and some text
		<p>and a paragraph
		<img src="./fake/image1.png">
		<div>
			finally a div with some
			<p>deeper nesting
			and another image for good measure
			<img src="./fake/image2.png">
		</div>
	</body>
</html>
`
	n, _ := html.Parse(strings.NewReader(s))
	words, images := countWordsAndImages(n)
	if words != wantedWords {
		t.Errorf("countWordsAndImages returned incorrect words %d, want %d", words, wantedWords)
	}
	if images != wantedImages {
		t.Errorf("countWordsAndImages returned incorrect images %d, want %d", images, wantedImages)
	}
}
