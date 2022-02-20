package main

import "fmt"
import "math"

func main() {
	var sitcomName string
	var episodeLength float64
	var brakeLenght float64
	fmt.Scanln(&sitcomName)
	fmt.Scanln(&episodeLength)
	fmt.Scanln(&brakeLenght)

	var restLength float64 = brakeLenght * 0.25
	var lunchLength float64 = brakeLenght * 0.125
	var availableTime float64 = brakeLenght - restLength - lunchLength  

	if availableTime >= episodeLength {
		var spareTime float64 = math.Ceil(availableTime - episodeLength)
		fmt.Printf("You have enough time to watch %s and left with %f minutes free time.", sitcomName, spareTime)
	} else {
		var neededTime float64 = math.Ceil(episodeLength - availableTime)
		fmt.Printf("You don't have enough time to watch %s, you need %f more minutes.", sitcomName, neededTime)
	}
}