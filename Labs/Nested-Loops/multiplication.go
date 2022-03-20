package main

import (
	"fmt"
)

func main() {
	// multiplier 1 and multiplier 2
	//
	for mul1 := 1; mul1 <= 10; mul1++ {
		for mul2 := 1; mul2 <= 10; mul2++ {
			fmt.Printf("%d * %d = %d", mul1, mul2, mul1 * mul2)
			fmt.Println()
		}
	}
}