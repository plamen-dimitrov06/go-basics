package main

import "fmt"

func main() {
	var points int
	var bonusPoints float32 = 0
	fmt.Scanln(&points)

	if points < 101 {
		bonusPoints = 5
	} else if points > 100 && points < 1001 {
		bonusPoints = float32(points) * 0.2
	} else if points > 1000 {
		bonusPoints = float32(points) * 0.1
	}

	if points % 2 == 0 {
		bonusPoints += 1
	} else if points % 5 == 0 && points % 2 != 0 {
		bonusPoints += 2
	}
	

	fmt.Println(bonusPoints)
	fmt.Println(float32(points) + bonusPoints)
}