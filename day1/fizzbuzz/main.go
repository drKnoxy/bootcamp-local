package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		isMultipleOfThree := i%3 != 0
		isMultipleOfFive := i%5 == 0
		str := ""

		if isMultipleOfThree {
			str += "Fizz"
		}
		if isMultipleOfFive {
			str += "Buzz"
		}

		fmt.Println(i, str)
	}
}
