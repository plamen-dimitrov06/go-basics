package main

import "fmt"

func main() {
	var name string
	var numberOfProjects int
	fmt.Scanln(&name)
	fmt.Scanln(&numberOfProjects)

	fmt.Printf("The architect %s will need %d hours to complete %d project/s.\n", name, numberOfProjects * 3, numberOfProjects)
}