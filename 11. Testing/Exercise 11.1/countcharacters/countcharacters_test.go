package countcharacters

import (
	"testing"
	"unicode/utf8"
)

// areCharacterCountsEqual returns true if all key/value pairs in left are in right and all key/value pairs in right is in left.
func areCharacterCountsEqual(left, right map[rune]int) bool {
	for r, count := range left {
		if count != right[r] {
			return false
		}
	}
	for r, count := range right {
		if count != left[r] {
			return false
		}
	}
	return true
}

func areUtflensEqual(left, right [utf8.UTFMax + 1]int) bool {

	for i := 1; i <= utf8.UTFMax; i++ {
		if left[i] != right[i] {
			return false
		}
	}
	return true
}

// areCategoryCountsEqual returns true if all key/value pairs in left are in right and all key/value pairs in right is in left.
func areCategoryCountsEqual(left, right map[string]int) bool {
	for r, count := range left {
		if count != right[r] {
			return false
		}
	}
	for r, count := range right {
		if count != left[r] {
			return false
		}
	}
	return true
}

func TestCountCharacters(t *testing.T) {
	s := `
1
22
333
4444
55555`
	characterCounts, utflen, invalid, categoryCounts := CountCharacters(s)

	wantedCharacterCounts := map[rune]int{
		'\n': 5,
		'1':  1,
		'2':  2,
		'3':  3,
		'4':  4,
		'5':  5,
	}
	var wantedUtflen [utf8.UTFMax + 1]int
	wantedUtflen[1] = utf8.RuneCountInString(s) // All characters are ASCII / 1-byte characters
	wantedInvalid := 0
	wantedCategoryCounts := map[string]int{
		"digit":      1 + 2 + 3 + 4 + 5,
		"whitespace": 5,
	}
	if !areCharacterCountsEqual(characterCounts, wantedCharacterCounts) {
		t.Error("CountCharacters returns incorrect character counts")
	}
	if !areUtflensEqual(utflen, wantedUtflen) {
		t.Error("CountCharacters returns incorrect utflen")
	}
	if invalid != wantedInvalid {
		t.Errorf("CountCharacters returns incorrect invalid %d, want %d", invalid, wantedInvalid)
	}
	if !areCategoryCountsEqual(categoryCounts, wantedCategoryCounts) {
		t.Error("CountCharacters returns incorrect category counts")
	}
}
