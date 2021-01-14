// Package intset provides a set of integers based on a bit vector.
package intset

import (
	"bytes"
	"fmt"
)

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// unsetFirstBit returns n with the first 1 bit set to 0.
// If n is equal to 0, 0 is returned.
func unsetFirstBit(n uint64) uint64 {
	return n & (n - 1)
}

// Len returns the number of elements in the set.
func (s *IntSet) Len() int {
	n := 0
	for _, word := range s.words {
		for ; word != 0; word = unsetFirstBit(word) {
			n++
		}
	}
	return n
}

// Remove removes x from the set.
// If x is not in the set it does nothing
func (s *IntSet) Remove(x int) {
	wordIndex, bitIndex := x/64, x%64
	if wordIndex < len(s.words) {
		s.words[wordIndex] = s.words[wordIndex] ^ (1 << bitIndex)
	}
}

// Copy returns a distinct copy of the set.
func (s *IntSet) Copy() *IntSet {
	sCopy := IntSet{
		words: make([]uint64, len(s.words)),
	}

	for i, word := range s.words {
		sCopy.words[i] = word
	}

	return &sCopy
}
