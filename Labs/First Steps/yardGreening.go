package main

import "fmt"

func main() {
	var squareFeetArea float32
	const discountPercent float32 = 0.18

	fmt.Scanln(&squareFeetArea)

	var subtotal float32 = squareFeetArea * 7.61
	var discount = subtotal * discountPercent
	var total float32 = subtotal - discount

	fmt.Printf("The final price is: %f lv.\n", total) 
	fmt.Printf("The discount is: %f lv.\n", discount) 
}