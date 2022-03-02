package main

import "fmt"

func main() {
	var flowerType string
	var flowersQuantity int
	var budget float32
	fmt.Scanln(&flowerType)
	fmt.Scanln(&flowersQuantity)
	fmt.Scanln(&budget)
	
	var total float32
	switch flowerType {
		case "Roses":
			total = float32(flowersQuantity) * 5
			if flowersQuantity > 80 {
				total = total * 0.9
			}
		case "Dahlias":
			total = float32(flowersQuantity) * 3.8
			if flowersQuantity > 90 {
				total = total * 0.85
			}
		case "Tulips":
			total = float32(flowersQuantity) * 2.8
			if flowersQuantity > 80 {
				total = total * 0.85
			}
		case "Narcissus":
			total = float32(flowersQuantity) * 3
			if flowersQuantity < 120 {
				total = total + (total * 0.15)
			}
		case "Gladiolus":
			total = float32(flowersQuantity) * 2.5
			if flowersQuantity < 80 {
				total = total + (total * 0.2)
			}
	}

	if total <= budget {
		fmt.Printf("Hey, you have a great garden with %d %s and %.2f leva left.", flowersQuantity, flowerType, budget - total)
	} else {
		fmt.Printf("Not enough money, you need %.2f leva more.", total - budget)
	}
}