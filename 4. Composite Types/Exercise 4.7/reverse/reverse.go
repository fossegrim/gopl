package reverse

// Reverse reverses the characters in an UTF-8 encoded string in place.
func Reverse(s []byte) {
	// reverse the bytes in each character of s, without changing the order of the characters.
	slen := len(s)
	for i, characterStart := 0, 0; i < slen; i++ {
		isLastByteInCharacter := i == slen - 1
		// if the next byte is the beginning of a character, the current is the end of one
		isLastByteInCharacter = isLastByteInCharacter || isBeginningOfCharacter(s[i+1])

		if isLastByteInCharacter {
			for j, k := characterStart, i; j <= k; j,k = j+1, k-1 {
				s[j], s[k] = s[k], s[j]
			}
			characterStart = i+1
		}
	}

	// reverse the bytes of s.
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// isBeginningOfCharacter returns true if the byte is the start of a multi-byte UTF-8 character or a ASCII character.
// Otherwise it returns false.
func isBeginningOfCharacter(c byte) bool {
	// If the most significant bits are NOT 0b10 the byte must be either an ASCII character or the start of a UTF-8 character.
	return (c & byte(0b11000000)) != byte(0b10000000)
}
