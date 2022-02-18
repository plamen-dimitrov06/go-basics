package main

import "fmt"

func main() {
	var firstSide int
	var secondSide int

	fmt.Scanln(&firstSide)
	fmt.Scanln(&secondSide)

	fmt.Println(firstSide * secondSide)
}