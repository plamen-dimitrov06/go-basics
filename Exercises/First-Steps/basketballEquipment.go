package main

import "fmt"

func main() {
	var yearlySubscription float32
	fmt.Scanln(&yearlySubscription)

	var basketballShoes = yearlySubscription * 0.6
	var basketballUniform = basketballShoes * 0.8
	var basketball = basketballUniform * 0.25
	var basketballAccessories = basketball * 0.2

	var total float32 = basketballShoes + basketballUniform + basketball + basketballAccessories

	fmt.Println(total + yearlySubscription)
}