package populate

import "golang.org/x/net/html"

// PopulateCountMap counts how many elements of name name there are in n for each key in countMap.
// It returns a map[string]int where each value corresponds to how many elements of name key was found.
func PopulateCountMap(n *html.Node, countMap map[string]int) map[string]int {
	for k, _ := range countMap {
		countMap[k] = countNames(n, k)
	}
	return countMap
}

// countNames counts how many elements of name name there are in n.
func countNames(n *html.Node, name string) (count int) {
	if n.Type == html.ElementNode && n.Data == name {
		count++
	}

	if n.FirstChild != nil {
		count += countNames(n.FirstChild, name)
	}
	if n.NextSibling != nil {
		count += countNames(n.NextSibling, name)
	}

	return
}
