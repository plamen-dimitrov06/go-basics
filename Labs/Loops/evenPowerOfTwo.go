package main

import "fmt"
import "math"

func main() {
	var numbersTotal int
	fmt.Scanln(&numbersTotal)

	fmt.Println(1)
	for i := 1; i <= numbersTotal; i++ {
		if i % 2 == 0 {
			fmt.Println(math.Pow(2, float64(i)))
		}
	}
}