package main

import (
	"fmt"
	"strconv"
)

func main() {
	var highest int = -10000
	for {
		var inputNumber string
		fmt.Scanln(&inputNumber)

		if inputNumber == "Stop" {
			fmt.Printf("%d", highest)
			break
		}

		if number, err := strconv.Atoi(inputNumber); err == nil {
			if number > highest {
				highest = number
			}
		}
	}
	
}