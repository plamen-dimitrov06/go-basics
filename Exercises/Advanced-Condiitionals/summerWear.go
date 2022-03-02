package main

import "fmt"

func main() {
	var tempature int
	var timeOfDay string
	fmt.Scanln(&tempature)
	fmt.Scanln(&timeOfDay)

	var outfit string
	var shoes string
	
	// perhaps the tempature should be checked before the day
	// since the day check is simpler
	if timeOfDay == "Morning" {
		switch true {
			case tempature >= 10 && tempature <= 18 :
				outfit = "Sweatshirt"
				shoes = "Sneakers"
			case tempature > 18 && tempature <= 24:
				outfit = "Shirt"
				shoes = "Moccasins"
			case tempature >= 25:
				outfit = "T-Shirt"
				shoes = "Sandals"
		}
	} else if timeOfDay == "Afternoon" {
		switch true {
			case tempature >= 10 && tempature <= 18:
				outfit = "Shirt"
				shoes = "Moccasins"
			case tempature > 18 && tempature <= 24:
				outfit = "T-Shirt"
				shoes = "Sandals"
			case tempature >= 25:
				outfit = "Swim Suit"
				shoes = "Barefoot"
		}
	} else {
		switch true {
			case tempature >= 10 && tempature <= 18:
				outfit = "Shirt"
				shoes = "Moccasins"
			case tempature > 18 && tempature <= 24:
				outfit = "Shirt"
				shoes = "Moccasins"
			case tempature >= 25:
				outfit = "Shirt"
				shoes = "Moccasins"
		}
	}

	fmt.Printf("It's %d degrees, get your %s and %s.", tempature, outfit, shoes)
}