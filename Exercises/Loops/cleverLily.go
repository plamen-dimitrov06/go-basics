package main

import "fmt"

func main() {
	var age int
	fmt.Scanln(&age)
	var washerPrice float32
	fmt.Scanln(&washerPrice)
	var toyPrice int
	fmt.Scanln(&toyPrice)

	var savedMoney = 0
	var toysCount = 0
	var moneyPerBirthday int = 10
	var counter float32 = 0
	for i := 1; i <= age; i++ {
		if i % 2 == 0 {
			savedMoney += moneyPerBirthday
			moneyPerBirthday += 10
			counter++
		} else {
			toysCount++
		}
	}

	var total float32 = float32(savedMoney + (toysCount * toyPrice))
	total = total - counter
	if total >= washerPrice {
		fmt.Printf("Yes! %.2f", total - washerPrice)
	} else {
		fmt.Printf("No! %.2f", washerPrice - total)
	}
}