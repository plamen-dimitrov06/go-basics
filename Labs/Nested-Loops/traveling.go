package main

import (
	"fmt"
	"strconv"
)

func main() {
	for {
		var country string
		fmt.Scanln(&country)
		if country == "End" {
			break
		}
		var budget float64
		fmt.Scanln(&budget)

		for i := budget; i > 0; {
			var savings string
			fmt.Scanln(&savings)
			if savings == "End" {
				break
			}
			if number, err := strconv.ParseFloat(savings, 32); err == nil {
				i-=number
			}
		}

		fmt.Printf("Going to %s! \n", country)
	}
}