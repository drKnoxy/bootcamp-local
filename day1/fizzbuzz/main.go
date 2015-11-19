package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		var isMultipleOfThree = i%3 != 0
		var isMultipleOfFive = i%5 == 0
		var isMultipleOfBoth = isMultipleOfThree && isMultipleOfFive
		var str string

		if isMultipleOfBoth {
			str = "FizzBuzz"
		} else if isMultipleOfThree {
			str = "Fizz"
		} else if isMultipleOfFive {
			str = "Buzz"
		}

		fmt.Println(i, str)
	}
}
