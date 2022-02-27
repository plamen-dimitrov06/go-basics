package main

import "fmt"

func main() {
	var age float32
	var gender string
	fmt.Scanln(&age)
	fmt.Scanln(&gender)

	if gender == "m" {
		if age >= 16 {
			fmt.Println("Mr.")
		} else {
			fmt.Println("Master")
		}
	} else {
		if age >= 16 {
			fmt.Println("Ms.")
		} else {
			fmt.Println("Miss")
		}
	}
}