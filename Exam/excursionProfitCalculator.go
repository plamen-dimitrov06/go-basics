package main

import (
	"fmt"
)

func main() {
	var mountainExcurions int
	fmt.Scanln(&mountainExcurions)
	var seaExcursions int
	fmt.Scanln(&seaExcursions)

	var vacation string
	fmt.Scanln(&vacation)

	var profit int
	for vacation != "Stop" {
		if vacation == "sea" && seaExcursions > 0 {
			profit += 680
			seaExcursions--
		} else if vacation == "mountain" && mountainExcurions > 0{
			profit += 499
			mountainExcurions--
		}

		if mountainExcurions <= 0 && seaExcursions <= 0 {
			break
		}

		fmt.Scanln(&vacation)
	}
	
	if mountainExcurions <= 0 && seaExcursions <= 0 {
		fmt.Println("Good job! Everything is sold.")
	}

	fmt.Printf("Profit: %d leva.", profit)
}