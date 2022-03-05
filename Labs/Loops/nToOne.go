package main

import "fmt"

func main() {
	var number int
	fmt.Scanln(&number)

	for counter := number; counter <= 1; counter-- {
		fmt.Println(counter)
	}
}