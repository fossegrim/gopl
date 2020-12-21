package comma

import (
	"bytes"
	"log"
)

// Comma inserts commas in the integral part of a number string.
// DONE: Handle signed numbers
// TODO: Handle decimal numbers
func Comma(number string) string {
	var buf bytes.Buffer
	var numberSansSign string
	switch number[0] {
	case '-', '+':
		buf.WriteByte(number[0])
		numberSansSign = number[1:]
	default:
		numberSansSign = number
	}

	// The number of unprocessed integral digits
	nIntegralDigits := 0
	isDecimalNumber := false
	for ; nIntegralDigits < len(numberSansSign); nIntegralDigits++ {
		if numberSansSign[nIntegralDigits] == '.' {
			isDecimalNumber = true
			break
		}
	}

	for i := 0; i < nIntegralDigits; i++ {
		digitsLeft := nIntegralDigits - i
		if digitsLeft%3 == 0 && digitsLeft != nIntegralDigits {
			err := buf.WriteByte(',')
			if err != nil {
				log.Fatal(err)
			}
		}
		err := buf.WriteByte(numberSansSign[i])
		if err != nil {
			log.Fatal(err)
		}
	}

	r := buf.String()
	if isDecimalNumber {
		r += "." + numberSansSign[nIntegralDigits+1:]
	}
	return r
}

// func Comma(number string) string {
// 	var buf bytes.Buffer
// 	var digits string
// 	switch number[0] {
// 	case '-', '+':
// 		buf.WriteByte(number[0])
// 		digits = number[1:]
// 	default:
// 		digits = number
// 	}

// 	ndigits := len(digits)
// 	for i := 0; i < ndigits; i++ {
// 		digitsLeft := ndigits - i
// 		if digitsLeft%3 == 0 && digitsLeft != ndigits {
// 			err := buf.WriteByte(',')
// 			if err != nil {
// 				log.Fatal(err)
// 			}
// 		}
// 		err := buf.WriteByte(digits[i])
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	}
// 	return buf.String()
// }
