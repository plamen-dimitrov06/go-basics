package main

import "fmt"

func main() {
	var firstNumber int
	var secondNumber int
	fmt.Scanln(&firstNumber)
	fmt.Scanln(&secondNumber)

	if firstNumber >= secondNumber  {
		fmt.Println(firstNumber)
	} else {
		fmt.Println(secondNumber)
	}
}