package main

import "fmt"

func main() {
	var goal int
	fmt.Scanln(&goal)

	var sum int
	for {
		var inputNumber int
		fmt.Scanln(&inputNumber)

		sum += inputNumber
		if sum < goal {
			continue
		} else {
			fmt.Printf("%d", sum)
			break
		}
	}
	
}