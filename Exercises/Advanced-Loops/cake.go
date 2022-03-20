package main

import (
	"fmt"
	"strconv"
	"math"
)

func main() {
	var widht float64
	fmt.Scanln(&widht)
	var height float64
	fmt.Scanln(&height)

	var cakeArea float64 = widht * height

	for {
		var piece string
		fmt.Scanln(&piece)

		if piece == "STOP" {
			fmt.Printf("%.0f pieces are left.", cakeArea)
			break
		}

		if number, err := strconv.ParseFloat(piece, 32); err == nil {
			cakeArea -= number
			if cakeArea < 0 {
				fmt.Printf("No more cake left! You need %.0f pieces more.", math.Abs(float64(cakeArea)))
				break
			}
		}
	}	
}