package main

import "fmt"

func main() {
	fmt.Println("Enter the amount of number you wanna add")
	var amountOfNumber int
	fmt.Scan(&amountOfNumber)

	var largestNumber int = 1
	var smallestNumber int

	for count := 1; count <= amountOfNumber; count++ {
		fmt.Println("Enter the amount of number you wanna add")
		var number int
		fmt.Scan(&number)

		if number > largestNumber {
			largestNumber = number
		}
		if number < largestNumber {
			smallestNumber = number
		}
	}

	fmt.Println("largest is ", largestNumber)
	fmt.Println("smallest is", smallestNumber)
}
