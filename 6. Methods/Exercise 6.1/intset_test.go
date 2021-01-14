package intset

import "testing"

func TestLen(t *testing.T) {
	wantedLen := 4

	var s IntSet
	s.Add(1)
	s.Add(3)
	s.Add(8)
	s.Add(9)
	len := s.Len()

	if len != wantedLen {
		t.Errorf("s.Len() returned %d, want %d", len, wantedLen)
	}
}

func TestLenEmpty(t *testing.T) {
	wantedLen := 0

	var s IntSet
	len := s.Len()

	if len != wantedLen {
		t.Errorf("s.Len() returned %d, want %d", len, wantedLen)
	}
}

func TestRemove(t *testing.T) {
	var s IntSet
	s.Add(4)
	s.Remove(4)

	if s.Has(4) != false {
		t.Errorf("s.Remove(4) does not remove 4 from s")
	}
}

// Test that attempting to remove 1 from an empty IntSet does not cause a panic.
// The reason that this could cause a panic, is that a naive implementation of s.Remove could try modifying the word which 1 would belong to if it was in the set, without checking that the slice even has the word which 1 would belong too.
// This would cause a index out of range panic.
func TestRemoveFromEmpty(t *testing.T) {
	var s IntSet
	s.Remove(1)

	if recover() != nil {
		t.Errorf("s.Remove(1) panicked")
	}
}

func TestCopy(t *testing.T) {
	var s IntSet
	intsInS := []int{5, 8, 74, 1, 76}
	for _, n := range intsInS {
		s.Add(n)
	}
	var c IntSet = *(s.Copy())

	if &s.words[0] == &c.words[0] {
		t.Errorf("s.Copy() returns a copy using the same underlying array as s")
	}

	if len(s.words) != len(c.words) {
		t.Errorf("s.Copy() does not return a copy with the same number of elements as s")
	}

	for _, n := range intsInS {
		if !c.Has(n) {
			t.Errorf("s.Copy(c) does not return a copy containing %d", n)
		}
	}
}
