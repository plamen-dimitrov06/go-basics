package main

import "fmt"

func main() {
	var packsOfDogFood float32
	var packsOfCatFood float32
	fmt.Scanln(&packsOfDogFood)
	fmt.Scanln(&packsOfCatFood)

	fmt.Printf("%f lv.\n", (packsOfDogFood * 2.5) + (packsOfCatFood * 4))
}