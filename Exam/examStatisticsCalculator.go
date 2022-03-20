package main

import (
	"fmt"
)

func main() {
	var studentsCount int
	fmt.Scanln(&studentsCount)

	var totalScoreSum float32

	var failScoreCount float32
	var poorScoreCount float32
	var averageScoreCount float32
	var excellentScoreCount float32

	for i := 0; i < studentsCount; i++ {
		var grade float32
		fmt.Scanln(&grade)

		totalScoreSum += grade

		switch {
			case grade < 3:
				failScoreCount++
			case grade >= 3 && grade <= 3.99:
				poorScoreCount++
			case grade >= 4 && grade <= 4.99 :
				averageScoreCount++
			case grade >= 5:
				excellentScoreCount++
		}
	}

	fmt.Printf("Top students: %.2f%% \n", (excellentScoreCount / float32(studentsCount)) * 100)
	fmt.Printf("Between 4.00 and 4.99: %.2f%% \n", (averageScoreCount / float32(studentsCount)) * 100)
	fmt.Printf("Between 3.00 and 3.99: %.2f%% \n", (poorScoreCount / float32(studentsCount)) * 100)
	fmt.Printf("Fail: %.2f%% \n", (failScoreCount / float32(studentsCount)) * 100)
	fmt.Printf("Average: %.2f", totalScoreSum / float32(studentsCount))
}