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
}
