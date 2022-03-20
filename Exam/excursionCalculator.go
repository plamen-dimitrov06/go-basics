package main

import (
	"fmt"
)

func main() {
	var peopleCount float32
	fmt.Scanln(&peopleCount)
	var season string
	fmt.Scanln(&season)

	var pricerPerPerson float32 = 0
	if peopleCount <= 5 {
		switch season {
			case "spring": pricerPerPerson = 50
			case "summer": pricerPerPerson = 48.5
			case "autumn": pricerPerPerson = 60
			case "winter": pricerPerPerson = 86
		}
	} else {
		switch season {
			case "spring": pricerPerPerson = 48
			case "summer": pricerPerPerson = 45
			case "autumn": pricerPerPerson = 49.5
			case "winter": pricerPerPerson = 85
		}
	}
	var total float32 = peopleCount * pricerPerPerson
	if season == "summer" {
		total = total * 0.85
	} else if season == "winter" {
		total = total + (total * 0.08)
	}

	fmt.Printf("%.2f leva.", total)
}