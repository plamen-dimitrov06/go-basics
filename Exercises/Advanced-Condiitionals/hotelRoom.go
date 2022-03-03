package main

import "fmt"

func main() {
	var month string
	var days float32
	fmt.Scanln(&month)
	fmt.Scanln(&days)

	var studioPrice float32
	var aparmentPrice float32
	var studioDiscount float32 = 1
	var apartmentDiscount float32 = 1
	if days > 14 {
		apartmentDiscount = 0.9
	}
	switch month {
		case "May", "October":
			switch {
				case days > 7 && days <= 14:
					studioDiscount = 0.95
				case days > 14:
					studioDiscount = 0.7
			}
			studioPrice = float32((50 * studioDiscount) * days)
			aparmentPrice = float32((65 * apartmentDiscount) * days)
		case "June", "September":
			if days > 14 {
				studioDiscount = 0.8
			}
			studioPrice = float32((75.2 * studioDiscount) * days)
			aparmentPrice = float32((68.7 * apartmentDiscount) * days)
		case "July", "August":
			studioPrice = float32(76 * days)
			aparmentPrice = float32((77 * apartmentDiscount) * days)
	}

	fmt.Printf("Apartment: %.2f lv.", aparmentPrice)
	fmt.Println()
	fmt.Printf("Studio: %.2f lv.", studioPrice)
}