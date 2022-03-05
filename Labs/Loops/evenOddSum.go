package main

import "fmt"
import "math"

func main() {
	var numbersTotal int
	fmt.Scanln(&numbersTotal)

	var evenSum int
	var oddSum int
	for i := 1; i <= numbersTotal; i++ {
		var number int
		fmt.Scanln(&number)
		if i % 2 == 0 {
			evenSum += number
		} else {
			oddSum += number
		}
	}

	if evenSum == oddSum {
		fmt.Println("Yes")
		fmt.Printf("Sum = %d", evenSum)
	} else {
		fmt.Println("No")
		fmt.Printf("Diff = %.0f", math.Abs(float64(evenSum - oddSum)))
	}
}