package main

import "fmt"

const milesPerKm = 1.60934

func main() {
	var mi float64
	var km float64

	fmt.Println("We're going to convert miles to kilometers.")
	fmt.Println("Enter miles: ")
	fmt.Scanf("%f", &mi)

	km = mi * milesPerKm
	fmt.Println("that's", km, "kilometers")
}
