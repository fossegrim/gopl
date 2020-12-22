package rotate

// Rotate rotates a n positions to the left in place.
// n must be positive.
func Rotate(a []int, n int) {
	alen := len(a)
	if n > alen {
		Rotate(a, n-alen)
		return
	}

	temp := make([]int, alen)
	copy(temp, a)

	copy(a[alen - n:], temp[:n])
	copy(a[:alen-n], temp[n:])
}
