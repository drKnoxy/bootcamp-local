package main

import "fmt"

const lb2kg = .453592

func main() {
	var lb float64
	var kg float64

	fmt.Println("We are going to convert Pounds to Kilograms!")
	fmt.Println("Enter a weight in pounds:")
	fmt.Scanf("%f", &lb)

	kg = lb * lb2kg

	fmt.Println("That's", kg, "kilograms.")
}
