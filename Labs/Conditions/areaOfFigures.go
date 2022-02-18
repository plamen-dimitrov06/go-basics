package main

import "fmt"
import "math"

func main() {
	var geometricalShape string
	fmt.Scanln(&geometricalShape)

	if geometricalShape == "square" {
		var side float32
		fmt.Scanln(&side)
		fmt.Printf("%.3f\n", side * side);
	} else if geometricalShape == "rectangle" {
		var firstSide float32
		var secondSide float32
		fmt.Scanln(&firstSide)
		fmt.Scanln(&secondSide)
		fmt.Printf("%.3f\n", firstSide * secondSide);
	} else if geometricalShape == "circle" {
		var radius float32
		fmt.Scanln(&radius)
		fmt.Printf("%.3f\n", math.Pi * (radius * radius))
	} else {
		var firstSide float32
		var secondSide float32
		fmt.Scanln(&firstSide)
		fmt.Scanln(&secondSide)
		fmt.Printf("%.3f\n", firstSide * secondSide / 2)
	}
	
}