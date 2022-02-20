package main

import "fmt"

func main() {
	var movieBudget float32
	var extras float32
	var extrasClothesPrice float32
	fmt.Scanln(&movieBudget)
	fmt.Scanln(&extras)
	fmt.Scanln(&extrasClothesPrice)

	// the decor is 10% of the budget
	//
	var decor float32 = movieBudget * 0.1

	var orderTotal float32 = extras * extrasClothesPrice

	// 10% discount if we need more than 150 extras
	//
	if extras > 150 {
		orderTotal = orderTotal * 0.9
	}

	var expenses float32 = decor + orderTotal

	if expenses > movieBudget {
		fmt.Println("Not enough money!")
		fmt.Printf("Wingard needs %.2f leva more.", expenses - movieBudget)
	} else {
		fmt.Println("Action!")
		fmt.Printf("Wingard start filming with %.2f leva left.", movieBudget - expenses)
	}
}