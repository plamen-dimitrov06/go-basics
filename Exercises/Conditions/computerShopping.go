package main

import "fmt"

func main() {
	var budget float32
	var graphicCards float32
	var cpus float32
	var memorySticks float32
	fmt.Scanln(&budget)
	fmt.Scanln(&graphicCards)
	fmt.Scanln(&cpus)
	fmt.Scanln(&memorySticks)

	var graphicCardsTotal float32 = graphicCards * 250
	var cpusTotal float32 = (graphicCardsTotal * 0.35) * cpus
	var memorySticksTotal float32 = (graphicCardsTotal * 0.1) * memorySticks
	var orderTotal float32 = graphicCardsTotal + cpusTotal + memorySticksTotal  

	// 10% discount if you order more GPUs than CPUs
	//
	if graphicCards > cpus {
		orderTotal = orderTotal * 0.85
	}

	if orderTotal <= budget {
		fmt.Printf("You have %.2f leva left!", budget - orderTotal)
	} else {
		fmt.Printf("Not enough money! You need %.2f leva more!", orderTotal - budget)
	}
}