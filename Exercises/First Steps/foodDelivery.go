package main

import "fmt"

func main() {
	var chickenMenus int
	var fishMenus int
	var vegetarianMenus int
	fmt.Scanln(&chickenMenus)
	fmt.Scanln(&fishMenus)
	fmt.Scanln(&vegetarianMenus)

	var chickenMenusTotal float32 = float32(chickenMenus) * 10.35
	var fishMenusTotal float32 = float32(fishMenus) * 12.4
	var vegetarianMenusTotal float32 = float32(vegetarianMenus) * 8.15

	var menusTotal float32 = chickenMenusTotal + fishMenusTotal + vegetarianMenusTotal
	var desert float32 = menusTotal * 0.2 
	const deliveryPrice = 2.5

	fmt.Println(menusTotal + desert + deliveryPrice)
}