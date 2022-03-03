package main

import "fmt"

func main() {
	var budget float32
	var season string
	fmt.Scanln(&budget)
	fmt.Scanln(&season)

	var total float32
	var area string
	var accommodation string = "Camp"
	if season != "summer" {
		accommodation = "Hotel"
	}
	if budget <= 100 {
		area = "Bulgaria"
		if season == "summer" {
			total = budget * 0.3	
		} else {
			total = budget * 0.7
		}
	} else if budget > 100 && budget <= 1000 {
		area = "Balkans"
		if season == "summer" {
			total = budget * 0.4	
		} else {
			total = budget * 0.8
		}
	} else {
		area = "Europe"
		total = budget * 0.9
		accommodation = "Hotel"
	}



	fmt.Printf("Somewhere in %s", area)
	fmt.Println()
	fmt.Printf("%s - %.2f", accommodation, total)
}