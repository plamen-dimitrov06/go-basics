package main

import (
	"fmt"
)

func main() {
	var start int
	fmt.Scanln(&start)
	var end int
	fmt.Scanln(&end)
	var magicNumber int
	fmt.Scanln(&magicNumber)

	var checksCounter = 0
	for j := start; j <= end; j++ {
		for i := start; i <= end; i++ {
			checksCounter++
			if i + j == magicNumber {
				fmt.Printf("Combination N:%d (%d + %d = %d)", checksCounter, j, i, magicNumber)
				return
			}
		}
	}

	fmt.Printf("%d combinations - neither equals %d", checksCounter, magicNumber)
}