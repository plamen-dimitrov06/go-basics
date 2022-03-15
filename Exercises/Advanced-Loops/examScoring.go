package main

import (
	"fmt"
)

func main() {
	var poorGradesLimit int
	fmt.Scanln(&poorGradesLimit)

	var scoresSum float32 = 0
	var examsTaken float32 = 0
	var lastProblem string
	var poorGradesCounter int = poorGradesLimit
	for {
		var subject string
		fmt.Scanln(&subject)

		if subject == "Enough" {
			fmt.Printf("Average score: %.2f", scoresSum / examsTaken)
			fmt.Println()
			fmt.Printf("Number of problems: %.0f", examsTaken)
			fmt.Println()
			fmt.Printf("Last problem: %s", lastProblem)
			fmt.Println()
			break
		}

		var grade float32
		fmt.Scanln(&grade)

		if grade <= 4 {
			poorGradesCounter--
			if poorGradesCounter == 0 {
				fmt.Printf("You need a break, %d poor grades.", poorGradesLimit)
				break
			}
		}
		
		lastProblem = subject
		examsTaken++
		scoresSum += grade
	}	
}