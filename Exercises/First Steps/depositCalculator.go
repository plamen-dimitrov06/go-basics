package main

import "fmt"

func main() {
	var initialAmount float32
	// the period below is in months
	//
	var depositPreiod float32
	var yearlyInterestRate float32
	fmt.Scanln(&initialAmount)
	fmt.Scanln(&depositPreiod)
	fmt.Scanln(&yearlyInterestRate)

	var interestAmount float32 = (initialAmount * (yearlyInterestRate/100))
	var monthlyInterestAmount float32 = (interestAmount / 12)

	fmt.Println(initialAmount + depositPreiod * monthlyInterestAmount)
}