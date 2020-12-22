package squashspace

import (
	"unicode"
	"unicode/utf8"
)
	
// Squash adjacent sequences of whitespace (as defined by unicode.IsSpace) into a single ASCII space.
// Squash requires the byteslice to be valid utf8 data.
func Squash(a *[]byte) {
	r, size := utf8.DecodeRune(*a)
	readIndex := size
	var writeIndex int
	if unicode.IsSpace(r) {
		(*a)[0] = ' '
		writeIndex = 1
	} else {
		writeIndex = size
	}

	prevRune := r
	for readIndex < len(*a) {
		r, size := utf8.DecodeRune((*a)[readIndex:])
		
		if unicode.IsSpace(r) {
			if !unicode.IsSpace(prevRune) {
				(*a)[writeIndex] = ' '
				writeIndex++
			}
		} else {
			// write character
			utf8.EncodeRune((*a)[writeIndex:], r)
			writeIndex += size
		}
		readIndex += size
		prevRune = r
	}
	*a = (*a)[:writeIndex]	// this changes our a, but not the a in the outer scope
}
