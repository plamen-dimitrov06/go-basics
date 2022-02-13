package main

import "fmt"

func main() {
	var leva float32
	fmt.Scanln(&leva)

	var UsdRate float32 = 1.79549
	fmt.Println(leva * UsdRate)
}