package main

import "fmt"

func main() {
	var lenght int
	var width int
	var height int
	var percent int // space filled with sand, pump etc.
	fmt.Scanln(&lenght)
	fmt.Scanln(&width)
	fmt.Scanln(&height)
	fmt.Scanln(&percent)

	var area int = lenght * width * height
	var liters float32 = float32(area) / 1000

	fmt.Println(liters * (1-float32(percent) / 100))
}