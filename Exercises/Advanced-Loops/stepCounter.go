package main

import (
	"fmt"
	"strconv"
)

func main() {
	var stepsCounter float64 = 0
	var isGoingHome bool = false
	for {
		var steps string
		fmt.Scanln(&steps)

		if steps == "GoingHome" {
			isGoingHome = true
			continue
		}

		if number, err := strconv.ParseFloat(steps, 32); err == nil {
			stepsCounter += number
			if stepsCounter >= 10000 {
				fmt.Println("Goal reached! Good job!")
				fmt.Printf("%.0f steps over the goal!", stepsCounter - 10000)
				break
			}
			if isGoingHome {
				fmt.Printf("%.0f more steps to reach goal.", 10000 - stepsCounter)
				break
			}
		}
	}	
}