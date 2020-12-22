package dedup

// Dedup removes adjacent duplicates from a in place.
func Dedup(a []string) {
	var writeIndex int
	for readIndex, writeIndex := 1, 1; readIndex < len(a); readIndex++ {
		if a[readIndex] != a[readIndex-1] {
			a[writeIndex] = a[readIndex]
			writeIndex++
		}
	}
	a = a[:writeIndex]
}
