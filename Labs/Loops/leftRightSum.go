package main

import "fmt"
import "math"

func main() {
	var numbersTotal int
	fmt.Scanln(&numbersTotal)

	var leftSum int
	for i := 0; i < numbersTotal; i++ {
		var number int
		fmt.Scanln(&number)
		leftSum += number
	}

	var rightSum int
	for i := 0; i < numbersTotal; i++ {
		var number int
		fmt.Scanln(&number)
		rightSum += number
	}

	if leftSum == rightSum {
		fmt.Printf("Yes, sum = %d", leftSum)
	} else {
		fmt.Printf("No, diff = %.0f", math.Abs(float64(leftSum - rightSum)))
	}
}