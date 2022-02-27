package main

import "fmt"

func main() {
	var vacationPrice float32
	var puzzles float32
	var dolls float32
	var bears float32
	var minions float32
	var trucks float32
	fmt.Scanln(&vacationPrice)
	fmt.Scanln(&puzzles)
	fmt.Scanln(&dolls)
	fmt.Scanln(&bears)
	fmt.Scanln(&minions)
	fmt.Scanln(&trucks)

	var totalQuantity = puzzles + dolls + bears + minions + trucks

	var orderTotal float32
	orderTotal = (puzzles * 2.6) + (dolls * 3) + (bears * 4.1) + (minions * 8.2) + (trucks * 2)

	// order 50 get 25% off
	//
	if totalQuantity >= 50 {
		orderTotal = orderTotal * 0.75
	}

	// 10% store rent 
	//
	orderTotal = orderTotal * 0.9

	if orderTotal >= vacationPrice {
		fmt.Printf("Yes! %.2f lv left.", orderTotal - vacationPrice)
	} else {
		fmt.Printf("Not enough money! %.2f lv needed.", vacationPrice - orderTotal)
	}
}