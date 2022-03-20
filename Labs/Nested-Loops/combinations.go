package main

import (
	"fmt"
)

func main() {
	var targetSum int
	fmt.Scanln(&targetSum)

	var matchesCounter = 0
	for num1 := 0; num1 <= targetSum; num1++ {
		for num2 := 0; num2 <= targetSum; num2++ {
			for num3 := 0; num3 <= targetSum; num3++ {
				if num1 + num2 + num3 == targetSum {
					matchesCounter++
				}
			}
		}
	}

	fmt.Println(matchesCounter)
}