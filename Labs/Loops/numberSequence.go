package main

import "fmt"

func main() {
	var numbersTotal int
	fmt.Scanln(&numbersTotal)

	var max int = -1
	var min int = 10000
	for i := 0; i < numbersTotal; i++ {
		var number int
		fmt.Scanln(&number)
		
		if number > max {
			max = number
		}

		if number < min {
			min = number
		}
	}

	if min == 10000 {
		min = max
	}

	fmt.Printf("Max number: %d", max)
	fmt.Println()
	fmt.Printf("Min number: %d", min)
}