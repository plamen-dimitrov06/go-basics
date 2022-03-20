package main

import (
	"fmt"
)

func main() {
	for hour := 0; hour < 24; hour++ {
		for minutes := 0; minutes < 60; minutes++ {
			fmt.Printf("%d:%d", hour, minutes)
			fmt.Println()
		}
	}
}