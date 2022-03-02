package main

import "fmt"

func main() {
	var budget float32
	var season string
	var fishers int
	fmt.Scanln(&budget)
	fmt.Scanln(&season)
	fmt.Scanln(&fishers)
	
	var rent float32
	switch season {
		case "Spring":
			rent = 3000
		case "Summer":
			rent = 4200
		case "Autumn":
			rent = 4200
		case "Winter":
			rent = 2600
	}

	switch {
		case fishers <= 6: rent = rent * 0.9
		case fishers > 6 && fishers <= 11: rent = rent * 0.85
		default: rent = rent * 0.75	
	}

	if fishers % 2 == 0 && season != "Autumn" {
		rent = rent * 0.95
	}

	if rent <= budget {
		fmt.Printf("Yes! You have %.2f leva left.", budget - rent)
	} else {
		fmt.Printf("Not enough money! You need %.2f leva.", rent - budget)
	}
}