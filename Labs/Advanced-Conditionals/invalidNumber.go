package main

import "fmt"

func main() {
	var number int
	fmt.Scanln(&number)

	if !(number >= 100 && number <= 200) && number != 0 {
		fmt.Println("invalid")
	}
}