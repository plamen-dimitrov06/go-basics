package main

import (
	"fmt"
)

func main() {
	var targetBook string
	fmt.Scanln(&targetBook)

	var booksChecked int = 0
	for {
		var book string
		fmt.Scanln(&book)

		if book == targetBook {
			fmt.Printf("You checked %d books and found it.", booksChecked)
			break
		}
		
		if book == "NoMoreBooks" {
			fmt.Println("The book you search is not here!")
			fmt.Printf("You checked %d books.", booksChecked)
			break
		}

		booksChecked++
	}	
}