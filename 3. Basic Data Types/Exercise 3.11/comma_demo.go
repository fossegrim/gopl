// Demonstrate comma package
package main

import (
	"fmt"

	"./comma"
)

func main() {
	fmt.Println(comma.Comma("1"))
	fmt.Println(comma.Comma("12"))
	fmt.Println(comma.Comma("123"))
	fmt.Println(comma.Comma("1234"))
	fmt.Println(comma.Comma("12345"))
	fmt.Println(comma.Comma("123456"))
	fmt.Println(comma.Comma("1234567"))
	fmt.Println(comma.Comma("12345678"))
	fmt.Println(comma.Comma("123456789"))
	println()

	fmt.Println(comma.Comma("1"))
	fmt.Println(comma.Comma("1.2"))
	fmt.Println(comma.Comma("12.3"))
	fmt.Println(comma.Comma("12.34"))
	fmt.Println(comma.Comma("123.45"))
	fmt.Println(comma.Comma("123.456"))
	fmt.Println(comma.Comma("1234.567"))
	fmt.Println(comma.Comma("1234.5678"))
	fmt.Println(comma.Comma("12345.6789"))
	fmt.Println(comma.Comma("12345.67899"))
	fmt.Println(comma.Comma("123456.78999"))
	fmt.Println(comma.Comma("123456.789999"))
	fmt.Println(comma.Comma("1234567.899999"))

	println()

	fmt.Println(comma.Comma("+1"))
	fmt.Println(comma.Comma("+12"))
	fmt.Println(comma.Comma("+123"))
	fmt.Println(comma.Comma("+1234"))
	fmt.Println(comma.Comma("+12345"))
	fmt.Println(comma.Comma("+123456"))
	fmt.Println(comma.Comma("+1234567"))
	fmt.Println(comma.Comma("+12345678"))
	fmt.Println(comma.Comma("+123456789"))
	println()

	fmt.Println(comma.Comma("+1"))
	fmt.Println(comma.Comma("+1.2"))
	fmt.Println(comma.Comma("+12.3"))
	fmt.Println(comma.Comma("+12.34"))
	fmt.Println(comma.Comma("+123.45"))
	fmt.Println(comma.Comma("+123.456"))
	fmt.Println(comma.Comma("+1234.567"))
	fmt.Println(comma.Comma("+1234.5678"))
	println()

	fmt.Println(comma.Comma("-1"))
	fmt.Println(comma.Comma("-1.2"))
	fmt.Println(comma.Comma("-12.3"))
	fmt.Println(comma.Comma("-12.34"))
	fmt.Println(comma.Comma("-123.45"))
	fmt.Println(comma.Comma("-123.456"))
	fmt.Println(comma.Comma("-1234.567"))
	fmt.Println(comma.Comma("-1234.5678"))
	fmt.Println(comma.Comma("-12345.6789"))
}
