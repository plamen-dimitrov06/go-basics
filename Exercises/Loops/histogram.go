package main

import "fmt"

func main() {
	var numbersTotal float32
	fmt.Scanln(&numbersTotal)

	var p1 float32 = 0
	var p2 float32 = 0
	var p3 float32 = 0
	var p4 float32 = 0
	var p5 float32 = 0
	for i := 0; i < int(numbersTotal); i++ {
		var number float32
		fmt.Scanln(&number)
		switch {
			case number < 200:
				p1++
			case number >= 200 && number < 400:
				p2++
			case number >= 400 && number < 600:
				p3++
			case number >= 600 && number < 800:
				p4++
			default :
				p5++
		}
	}

	fmt.Printf("%.2f%%", float32((p1 / numbersTotal) * 100));
	fmt.Println()
	fmt.Printf("%.2f%%", float32((p2 / numbersTotal) * 100));
	fmt.Println()
	fmt.Printf("%.2f%%", float32((p3 / numbersTotal) * 100));
	fmt.Println()
	fmt.Printf("%.2f%%", float32((p4 / numbersTotal) * 100));
	fmt.Println()
	fmt.Printf("%.2f%%", float32((p5 / numbersTotal) * 100));
	fmt.Println()
}