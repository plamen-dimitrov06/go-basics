package main

import "fmt"
import "math"

func main() {
	var currentRecord float64
	var distance float64
	var timePerMeter float64
	fmt.Scanln(&currentRecord)
	fmt.Scanln(&distance)
	fmt.Scanln(&timePerMeter)

	var newTime float64 = distance * timePerMeter
	var deduction = math.Floor(distance / 15)
	var timeDeduction = float64(deduction) * 12.5

	newTime += timeDeduction

	if currentRecord > newTime {
		fmt.Printf("Yes, he succeeded! The new world record is %.2f seconds.", newTime)
	} else {
		fmt.Printf("No, he failed! He was %.2f seconds slower.", newTime - currentRecord)
	}
}