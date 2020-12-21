package comma

import (
	"bytes"
	"log"
)

// Comma inserts commas in a non-negative decimal integer string.
func Comma(s string) string {
	var buf bytes.Buffer
	lastDigitIndex := len(s) - 1
	for digitsLeft := lastDigitIndex; digitsLeft >= 0; digitsLeft-- {
		err := buf.WriteByte(s[lastDigitIndex-digitsLeft])
		if err != nil {
			log.Fatal(err)
		}
		if digitsLeft%3 == 0 && digitsLeft != 0 {
			err := buf.WriteByte(',')
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	return buf.String()
}
