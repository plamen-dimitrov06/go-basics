package main

import "fmt"

func main() {
	var tenisMatches float32
	fmt.Scanln(&tenisMatches)
	var currentPoints int
	fmt.Scanln(&currentPoints)

	var totalPointsEarned int
	var winsCount float32
	for i := 0; i < int(tenisMatches); i++ {
		var position string
		fmt.Scanln(&position)
		if position == "F" {
			totalPointsEarned += 1200
		} else if position == "SF" {
			totalPointsEarned += 720
		} else if position == "W" {
			winsCount++
			totalPointsEarned += 2000
		}
	}

	fmt.Printf("Final points: %d", currentPoints + totalPointsEarned)
	fmt.Println()
	fmt.Printf("Average points: %d", totalPointsEarned / int(tenisMatches))
	fmt.Println()
	fmt.Printf("%.2f%%", float32((winsCount / tenisMatches) * 100))
	fmt.Println()

}