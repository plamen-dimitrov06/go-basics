package main

import "fmt"

func main() {
	for {
		var line string
		fmt.Scanln(&line)

		if line != "Stop" {
			fmt.Println(line)
		} else {
			break
		}
	}
	
}