package main

import (
	"fmt"
)

func main() {
	var vacationPrice float32
	fmt.Scanln(&vacationPrice)
	var availableBudget float32
	fmt.Scanln(&availableBudget)

	var spendCounter int = 0
	var daysCounter int = 0
	for {
		var action string
		fmt.Scanln(&action)
		var amount float32
		fmt.Scanln(&amount)
		daysCounter++
		if action == "spend" {
			spendCounter++
			if spendCounter == 5 {
				fmt.Println("You can't save the money.")
				fmt.Printf("%d", daysCounter)
				break
			}
			availableBudget -= amount
			if availableBudget < 0 {
				availableBudget = 0
			}
		} else {
			spendCounter = 0
			availableBudget += amount
			if availableBudget >= vacationPrice {
				fmt.Printf("You saved the money for %d days.", daysCounter)
				break
			}
		}

	
	}	
}