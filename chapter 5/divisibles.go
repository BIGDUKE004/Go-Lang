package main

import "fmt"

func main(){
	var sum int
	for count := 1; count <= 30; count++{
		if count % 3 == 0{
			sum += count
		}
	}

	fmt.Println(sum)
}
