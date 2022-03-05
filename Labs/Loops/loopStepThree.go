package main

import "fmt"

func main() {
	var number int
	fmt.Scanln(&number)

	for i := 1; i <= number; i+=3 {
		fmt.Println(i)
	}
}