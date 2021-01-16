package intset

import (
	"testing"
)

func TestAddAll(t *testing.T) {
	wantedLen := 3

	var s IntSet
	s.AddAll(1, 2, 3)
	len := s.Len()

	if len != wantedLen {
		t.Errorf("s.AddAll(1,2,3) does not add three elements to s")
	}
}
