package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		str := "odd"
		if i%2 == 0 {
			str = "even"
		}

		fmt.Println(i, str)
	}
}
