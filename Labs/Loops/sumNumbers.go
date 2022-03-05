package main

import "fmt"

func main() {
	var numbersTotal int
	fmt.Scanln(&numbersTotal)

	var sum int
	for i := 0; i < numbersTotal; i++ {
		var number int
		fmt.Scanln(&number)
		sum += number
	}

	fmt.Println(sum)
}