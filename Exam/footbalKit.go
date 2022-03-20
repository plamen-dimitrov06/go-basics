package main

import (
	"fmt"
)

func main() {
	var shirtPrice float32
	fmt.Scanln(&shirtPrice)
	var priceTarget float32
	fmt.Scanln(&priceTarget)

	var shortsPrice float32 = shirtPrice * 0.75
	var socksPrice float32 = shortsPrice * 0.2
	var shoesPrice float32 = 2 * (shortsPrice + shirtPrice)
	var grandTotal float32 = (shirtPrice + shortsPrice + socksPrice + shoesPrice) * 0.85

	if grandTotal >= priceTarget {
		fmt.Println("Yes, he will earn the world-cup replica ball!")
		fmt.Printf("His sum is %.2f lv.", grandTotal)
	} else {
		fmt.Println("No, he will not earn the world-cup replica ball.")
		fmt.Printf("He needs %.2f lv. more.", priceTarget - grandTotal)
	}
}