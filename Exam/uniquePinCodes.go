package main

import (
	"fmt"
)

func main() {
	var first int
	fmt.Scanln(&first)
	var second int
	fmt.Scanln(&second)
	var third int
	fmt.Scanln(&third)

	for digit1 := 2; digit1 <= first; digit1++ {
		for digit2 := 2; digit2 <= second; digit2++ {
			for digit3 := 2; digit3 <= third; digit3++ {
				if digit1 % 2 == 0 && digit3 % 2 == 0 && digit2 != 6 && digit2 != 4 && digit2 < 8{
					fmt.Println(digit1, digit2, digit3)
				}
			}
		}
	}

}