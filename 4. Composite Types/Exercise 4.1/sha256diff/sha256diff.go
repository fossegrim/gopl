package sha256diff

// Count counts the number of differing bits in two sha256 hashes
func Count(leftHash, rightHash [32]uint8) int {
	count := 0
	for i, leftByte := range leftHash {
		rightByte := rightHash[i]
		diff := leftByte ^ rightByte
		for diff != 0 {
			count++
			diff = diff & (diff - 1) // clear the rightmost 1-bit
		}
	}
	return count
}
