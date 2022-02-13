package main

import "fmt"

func main() {
	var maskingDrape int // sqr meter
	var paint float32 // in liters
	var paintDisolver int 	// in liters
	var paintersTime int
	fmt.Scanln(&maskingDrape)
	fmt.Scanln(&paint)
	fmt.Scanln(&paintDisolver)
	fmt.Scanln(&paintersTime)

	var drapeTotal float32 = float32(maskingDrape + 2) * 1.5
	var paintTotal float32 = (paint + (paint * 0.1)) * 14.5
	var disolverTotal float32 = float32(paintDisolver) * 5
	
	var total float32 = drapeTotal + paintTotal + disolverTotal + 0.4
	var laborCost float32 = (total * 0.3) * float32(paintersTime)
	
	fmt.Println(total + laborCost)
}