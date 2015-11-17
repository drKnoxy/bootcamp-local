package main

import "fmt"

func main() {
	var f float64
	var c float64

	fmt.Println("We are going to convert Fahrenheit to Celsius!")
	fmt.Println("Enter a temp in F:")
	fmt.Scanf("%f", &f)

	c = (f - 32) * 5 / 9

	fmt.Println("That's", c, "degrees celsius.")
}
