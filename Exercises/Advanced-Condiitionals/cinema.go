package main

import "fmt"

func main() {
	var projectionType string
	var rows int
	var columns int
	fmt.Scanln(&projectionType)
	fmt.Scanln(&rows)
	fmt.Scanln(&columns)

	var seats int = rows * columns

	switch projectionType {
		case "Premiere":
			fmt.Printf("%.2f leva", float32(seats) * 12)
		case "Normal":
			fmt.Printf("%.2f leva", float32(seats) * 7.5)
		case "Discount":
			fmt.Printf("%.2f leva", float32(seats) * 5)
	}
}