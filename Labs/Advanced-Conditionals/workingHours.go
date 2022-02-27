package main

import "fmt"

func main() {
	var hour int
	var day string
	fmt.Scanln(&hour)
	fmt.Scanln(&day)

	if hour >= 10 && hour <= 18 && day != "Sunday" {
		fmt.Println("open")
	} else {
		fmt.Println("closed")
	}
}