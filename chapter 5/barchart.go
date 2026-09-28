package main

import "fmt"

func main() {
	fmt.Println("Enter a number: ")
	var input int
	fmt.Scan(&input)

	var singleDigit int
	
	for counter := 1; counter <= input; counter++ {
		singleDigit = input % 10
		input = input / 10

	}

	for count := 0; count <= singleDigit; count++ {
		fmt.Println("*")
	}
}
