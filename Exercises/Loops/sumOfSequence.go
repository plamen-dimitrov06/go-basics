package main

import "fmt"
import "math"

func main() {
	var numbersTotal int
	fmt.Scanln(&numbersTotal)

	var max int = -10000
	var sum int = 0
	for i := 0; i < numbersTotal; i++ {
		var number int
		fmt.Scanln(&number)
		if number > max {
			max = number
		}
		sum += number
	}

	if sum - max == max {
		fmt.Println("Yes")
		fmt.Printf("Sum = %d", max)
	} else {
		fmt.Println("No")
		fmt.Printf("Diff = %.0f", math.Abs(float64(max - (sum - max))))
	}
}