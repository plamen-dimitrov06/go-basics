package main

import "fmt"
import "math"

func main() {
	var examHours int
	var examMinutes int
	var arrivalHours int
	var arrivalMinutes int
	fmt.Scanln(&examHours)
	fmt.Scanln(&examMinutes)
	fmt.Scanln(&arrivalHours)
	fmt.Scanln(&arrivalMinutes)

	examMinutes = examMinutes + (examHours * 60)
	arrivalMinutes = arrivalMinutes + (arrivalHours * 60)

	var beforeAfter string
	if arrivalMinutes > examMinutes {
		fmt.Println("Late")
		beforeAfter = "after"
	} else if (examMinutes == arrivalMinutes) || (examMinutes - arrivalMinutes) <= 30 {
		fmt.Println("On time")
		if examMinutes == arrivalMinutes {
			return
		}
		beforeAfter = "before"
	} else if examMinutes - arrivalMinutes > 30 {
		fmt.Println("Early")
		beforeAfter = "before"
	}
	var spareMinutes int
	if beforeAfter == "after" {
		spareMinutes = arrivalMinutes - examMinutes
	} else {
		spareMinutes = examMinutes - arrivalMinutes
	}

	var hoursEarly = math.Floor(float64(spareMinutes / 60))
	var minutesEarly = spareMinutes % 60
	if hoursEarly > 0 {
		fmt.Printf("%.0f:%.2d hours %s the start", hoursEarly, minutesEarly, beforeAfter)
	} else {
		fmt.Printf("%.0d minutes %s the start", minutesEarly, beforeAfter)
	}
}