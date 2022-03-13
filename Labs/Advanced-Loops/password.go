package main

import "fmt"

func main() {
	var user string
	var pass string
	fmt.Scanln(&user)
	fmt.Scanln(&pass)

	for {
		var inputPass string
		fmt.Scanln(&inputPass)
		if inputPass != pass {
			continue
		} else {
			fmt.Printf("Welcome %s!", user)
			break
		}
	}
	
}