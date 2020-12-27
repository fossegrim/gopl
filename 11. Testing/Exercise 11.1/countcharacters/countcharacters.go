package countcharacters

import (
	"unicode"
	"unicode/utf8"
)

// countCharacters counts the following things about the characters in s.
// - How many there are of each character.
// - How many there are of each length(number of bytes).
// - How many there are which are invalid.
// - How many there are which belongs to the categories "digit", "letter" and "whitespace".
func CountCharacters(s string) (map[rune]int, [utf8.UTFMax + 1]int, int, map[string]int) {
	characterCounts := make(map[rune]int)  // counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int        // count of lengths of UTF-8 encodings
	invalid := 0                           // count of invalid UTF-8 characters
	categoryCounts := make(map[string]int) // counts of Unicode characters belonging to specific categories

	for _, r := range s {
		n := utf8.RuneLen(r)
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		characterCounts[r]++
		utflen[n]++
		if unicode.IsDigit(r) {
			categoryCounts["digit"]++
		}
		if unicode.IsLetter(r) {
			categoryCounts["letter"]++
		}
		if unicode.IsSpace(r) {
			categoryCounts["whitespace"]++
		}
	}

	return characterCounts, utflen, invalid, categoryCounts
}
