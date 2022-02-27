package main

import "fmt"
import "math"

func main() {
	var number float64
	fmt.Scanln(&number)

	if math.Abs(number) <= 100 && number != 0 {
		fmt.Println("Yes")
	} else { 
		fmt.Println("No")
	}
}