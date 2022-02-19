package main

import "fmt"

func main() {
	var firstCompetitorTime int
	var secondCompetitorTime int
	var thirdCompetitorTime int
	fmt.Scanln(&firstCompetitorTime)
	fmt.Scanln(&secondCompetitorTime)
	fmt.Scanln(&thirdCompetitorTime)

	var secondsTotal int = firstCompetitorTime + secondCompetitorTime + thirdCompetitorTime
	
	var minutes int = secondsTotal / 60
	var seconds int = secondsTotal % 60
	
	fmt.Printf("%d:%.2d", minutes, seconds)
}