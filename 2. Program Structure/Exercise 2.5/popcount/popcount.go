package popcount

// pc[i] is the populcation count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
		//      ^         ^
		//      |         |
		//      |         value of the rightmost bit
		//      |
		//      population count of digits prior to i
	}
}

// FlatPopCount returns the population count (number of set bits) of x.
func FlatPopCount(x uint64) int {
	return int(0 +
		pc[byte(x>>(0*8))] + // population count of first 8 bits
		pc[byte(x>>(1*8))] + // population count of the next 8 bits
		pc[byte(x>>(2*8))] + // ...
		pc[byte(x>>(3*8))] + // ...
		pc[byte(x>>(4*8))] + // ...
		pc[byte(x>>(5*8))] + // ...
		pc[byte(x>>(6*8))] + // ...
		pc[byte(x>>(7*8))]) // population count of the final 8 bits
}

// LoopedPopCount returns the population count (number of set bits) of x.
func LoopedPopCount(x uint64) int {
	populationCount := 0
	for i := 0; i < 8; i++ {
		populationCount += int(pc[byte(x>>(i*8))])
	}
	return populationCount
}

// NoTablePopCount returns the population count (number of set bits)
func NoTablePopCount(x uint64) int {
	populationCount := 0
	for i := 0; i < 64; i++ {
		populationCount += int((x >> i) & 1)
	}
	return populationCount
}

// EpicPopCount returns the population count (number of set bits)
func EpicPopCount(x uint64) int {
	populationCount := 0
	for x != 0 {
		x = x & (x - 1) // clear the rightmost non-zero digit
		populationCount++
	}
	return populationCount
}
