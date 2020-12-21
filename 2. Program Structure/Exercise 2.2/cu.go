// Convert units
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"./lengthconv"
	"./tempconv"
	"./weightconv"
)

func main() {
	parseFloat := func(s string) float64 {
		x, err := strconv.ParseFloat(s, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		return x
	}

	numbers := os.Args[1:]
	if len(numbers) == 0 {
		input := bufio.NewScanner(os.Stdin)
		// This reads line by line, which is much inferior to word* by word, but the book has not yet showed any methods, direct** or indirect***, of doing this.
		// *by word I mean a whitespace separated sequence of characters
		// **a direct method would be a function, or the-like, which reads a word.
		// ***an indirect method would be a way that is not direct**. For example by reading a line, or so, and parsing that. Unfortunately, not enough string manipulation is taught to do so. strings.{Split,Join} is not enough
		for input.Scan() {
			x := parseFloat(input.Text())
			printConversions(x)
		}
	} else {
		for _, arg := range numbers {
			x := parseFloat(arg)
			printConversions(x)
		}
	}
}

func printConversions(x float64) {
	printTemperatureConversions(x)
	printLengthConversions(x)
	printWeightConversions(x)
	fmt.Println()
}

func printTemperatureConversions(x float64) {
	c := tempconv.Celsius(x)
	fmt.Printf("%v = %v = %v, ", c, tempconv.CToF(c), tempconv.CToK(c))

	f := tempconv.Fahrenheit(x)
	fmt.Printf("%v = %v = %v, ", f, tempconv.FToC(f), tempconv.FToK(f))

	k := tempconv.Kelvin(x)
	fmt.Printf("%v = %v = %v\n", k, tempconv.KToC(k), tempconv.KToF(k))
}

func printLengthConversions(x float64) {
	f := lengthconv.Feet(x)
	fmt.Printf("%v = %v = %v, ", f, lengthconv.FToI(f), lengthconv.FToM(f))

	i := lengthconv.Inch(x)
	fmt.Printf("%v = %v = %v, ", i, lengthconv.IToF(i), lengthconv.IToM(i))

	m := lengthconv.Meter(x)
	fmt.Printf("%v = %v = %v\n", m, lengthconv.MToI(m), lengthconv.MToF(m))
}

func printWeightConversions(x float64) {
	kg := weightconv.Kilogram(x)
	fmt.Printf("%v = %v, ", kg, weightconv.KToP(kg))

	lb := weightconv.Pound(x)
	fmt.Printf("%v = %v\n", lb, weightconv.PToK(lb))
}
