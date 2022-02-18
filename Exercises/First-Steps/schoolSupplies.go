package main

import "fmt"

func main() {
	var penPacks int
	var markerPacks int
	var markerCleaner int 	// in liters
	var discount float32
	fmt.Scanln(&penPacks)
	fmt.Scanln(&markerPacks)
	fmt.Scanln(&markerCleaner)
	fmt.Scanln(&discount)

	var pensTotal float32 = float32(penPacks) * 5.8
	var markersTotal float32 = float32(markerPacks) * 7.2
	var cleanerTotal float32 = float32(markerCleaner) * 1.2

	var total float32 = (pensTotal + markersTotal + cleanerTotal) * ((100 - discount) / 100)

	fmt.Println(total)
}