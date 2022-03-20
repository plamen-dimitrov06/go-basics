package main

import (
	"fmt"
	"strconv"
)

func main() {
	var sum float64 = 0
	for {
		var inputNumber string
		fmt.Scanln(&inputNumber)

		if inputNumber == "NoMoreMoney" {
			fmt.Printf("Total: %.2f", sum)
			break
		}
		if number, err := strconv.ParseFloat(inputNumber, 32); err == nil {
			if number < 0 {
				fmt.Println("Invalid operation!")
				fmt.Printf("Total: %.2f", sum)
				break
			} else {
				sum += number
				fmt.Printf("Increase: %.2f", number)
				fmt.Println();
			}
		}
	}
	
}