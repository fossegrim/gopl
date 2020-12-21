package main

import (
	"fmt"

	"./tempconv"
)

func main() {
	temperatureReport(0)
	temperatureReport(123)
	temperatureReport(-273.15)
}

func temperatureReport(temperatureC tempconv.Celsius) {
	var temperatureF tempconv.Fahrenheit = tempconv.CToF(temperatureC)
	var temperatureK tempconv.Kelvin = tempconv.FToK(temperatureF)

	fmt.Println("---===---")
	fmt.Printf("Celsius: %v\n", temperatureC)
	fmt.Printf("Fahrenheit: %v\n", temperatureF)
	fmt.Printf("Kelvin: %v\n", temperatureK)
	fmt.Println("---===---")
}
