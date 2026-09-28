package main

import "fmt"

func main() {
	fmt.Println("Enter Number")
	var number int
	fmt.Scan(&number)

	var sum int

	fmt.Printf("%s  %s \n", "n", "sum")

	for count := 1; count <= number; count++ {
		for counter := 1; counter <= count; counter++ {
			sum += counter
		}
		fmt.Printf("%d  %d \n", count, sum)
	}
}
