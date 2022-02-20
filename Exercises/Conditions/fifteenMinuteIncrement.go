package main

import "fmt"

func main() {
	var hours int
	var minutes int
	fmt.Scanln(&hours)
	fmt.Scanln(&minutes)

	minutes += 15

	if minutes >= 60 {
		hours += 1
		minutes = minutes % 60
	}

	if hours >= 24 {
		hours = 0
	}
	
	fmt.Printf("%d:%.2d", hours, minutes)
}