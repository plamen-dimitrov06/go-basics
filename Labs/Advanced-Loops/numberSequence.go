package main

import "fmt"

func main() {
	var goal int
	fmt.Scanln(&goal)

	var sum int = 1
	fmt.Println(sum)
	for {
		sum = sum * 2 + 1
		if sum <= goal {
			fmt.Println(sum)
			continue
		} else {
			break
		}
	}
	
}