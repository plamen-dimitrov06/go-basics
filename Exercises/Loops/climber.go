package main

import "fmt"

func main() {
	var groups float32
	fmt.Scanln(&groups)

	var musala float32 = 0
	var monblan float32 = 0
	var kilimanjaro float32 = 0
	var k2 float32 = 0
	var everest float32 = 0
	var totalClimbers float32 = 0
	for i := 0; i < int(groups); i++ {
		var numberOfPeople float32
		fmt.Scanln(&numberOfPeople)
		totalClimbers += numberOfPeople
		switch {
			case numberOfPeople <= 5:
				musala += numberOfPeople
			case numberOfPeople >= 6 && numberOfPeople <= 12:
				monblan += numberOfPeople
			case numberOfPeople >= 13 && numberOfPeople <= 25:
				kilimanjaro += numberOfPeople
			case numberOfPeople >= 26 && numberOfPeople <= 40:
				k2 += numberOfPeople
			default :
				everest += numberOfPeople
		}
	}

	fmt.Printf("%.2f%%", float32((musala / totalClimbers) * 100));
	fmt.Println()
	fmt.Printf("%.2f%%", float32((monblan / totalClimbers) * 100));
	fmt.Println()
	fmt.Printf("%.2f%%", float32((kilimanjaro / totalClimbers) * 100));
	fmt.Println()
	fmt.Printf("%.2f%%", float32((k2 / totalClimbers) * 100));
	fmt.Println()
	fmt.Printf("%.2f%%", float32((everest / totalClimbers) * 100));
	fmt.Println()
}