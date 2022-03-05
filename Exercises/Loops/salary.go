package main

import "fmt"

func main() {
	var tabs int
	fmt.Scanln(&tabs)
	var salary int
	fmt.Scanln(&salary)

	for i := 0; i < tabs; i++ {
		var website string
		fmt.Scanln(&website)
		if website == "Facebook" {
			salary -= 150
		} else if website == "Instagram" {
			salary -= 100
		} else if website == "Reddit" {
			salary -= 50
		}

		if salary <= 0 {
			fmt.Println("You have lost your salary.")
			return
		}
	}

	fmt.Println(salary)
}