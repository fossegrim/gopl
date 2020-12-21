package anagram

// IsAnagram returns true if s1 and s2 are anagrams of each other, otherwise it returns false.
// It handles non-ascii unicode characters correctly.
func IsAnagram(s1, s2 string) bool {
	s1CharacterCount, s2CharacterCount := countCharacters(s1), countCharacters(s2)

	// Verify that all characters in s1 appears in s2 the same number of times
	for character, count := range s1CharacterCount {
		if s2CharacterCount[character] != count {
			return false
		}
	}
	// Verify that all characters in s2 appears in s1 the same number of times
	for character, count := range s2CharacterCount {
		if s1CharacterCount[character] != count {
			return false
		}
	}

	return true
}

func countCharacters(s string) map[rune]int {
	characterCount := make(map[rune]int)
	for _, character := range s {
		characterCount[character]++
	}
	return characterCount
}
