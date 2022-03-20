package main

import (
	"fmt"
)

func main() {
	var people float32
	fmt.Scanln(&people)
	var days float32
	fmt.Scanln(&days)
	var transportCards float32
	fmt.Scanln(&transportCards)
	var museumTickets float32
	fmt.Scanln(&museumTickets)

	var daysCost float32 = days * 20
	var cardsCost float32 = transportCards * 1.6
	var museumCost float32 = museumTickets * 6
	var costPerPerson = daysCost + cardsCost + museumCost
	var groupCost float32 = people * costPerPerson

	var total = groupCost + (groupCost * 0.25)
	fmt.Printf("%.2f", total)
}