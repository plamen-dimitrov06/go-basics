package main

import (
	"fmt"
	"strconv"
)

func main() {
	var lowest int = 10000
	for {
		var inputNumber string
		fmt.Scanln(&inputNumber)

		if inputNumber == "Stop" {
			fmt.Printf("%d", lowest)
			break
		}

		if number, err := strconv.Atoi(inputNumber); err == nil {
			if number < lowest {
				lowest = number
			}
		}
	}
	
}