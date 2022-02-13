package main

import "fmt"
import "math"

func main() {
	var radians float32
	fmt.Scanln(&radians)

	fmt.Println(radians * 180 / math.Pi)
}