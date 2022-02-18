package main

import "fmt"

func main() {
	var inches float32
	fmt.Scanln(&inches)

	var oneInchInCm float32 = 2.54
	fmt.Println(inches * oneInchInCm)
}