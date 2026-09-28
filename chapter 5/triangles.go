package main

import "fmt"

func main() {
	for count := 1; count <= 10; count++ {
		for counter := 1; counter <= count; counter++ {
			fmt.Print("x")
		}
		fmt.Println(" ")
	}

	for count := 10; count >= 1; count-- {
		for counter := count; counter >= 1; counter-- {
			fmt.Print("x")
		}
		fmt.Println(" ")
	}

	for count := 1; count <= 10; count++ {
		for space := 1; space <= count; space++ {
			fmt.Print(" ")
		}

		for counter := 10; counter >= count; counter-- {
			fmt.Println("x")
		}
		fmt.Print()
	}
}
