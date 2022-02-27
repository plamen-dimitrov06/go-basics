package main

import "fmt"

func main() {
	var town string
	var sales float32
	fmt.Scanln(&town)
	fmt.Scanln(&sales)

	if sales < 0 {
		fmt.Println("error")
		return
	}

	if town == "Varna" {
		switch true {
			case sales >= 0 && sales <= 500: fmt.Printf("%.2f", sales * 0.045)
			case sales > 500 && sales <= 1000: fmt.Printf("%.2f", sales * 0.075)
			case sales > 1000 && sales <= 10000: fmt.Printf("%.2f", sales * 0.1)
			case sales > 10000: fmt.Printf("%.2f", sales * 0.13)
		}
	} else if town == "Sofia" {
		switch true {
			case sales >= 0 && sales <= 500: fmt.Printf("%.2f", sales * 0.05)
			case sales > 500 && sales <= 1000: fmt.Printf("%.2f", sales * 0.07)
			case sales > 1000 && sales <= 10000: fmt.Printf("%.2f", sales * 0.08)
			case sales > 10000: fmt.Printf("%.2f", sales * 0.12)
		}
	} else if town == "Plovdiv" {
		switch true {
			case sales >= 0 && sales <= 500: fmt.Printf("%.2f", sales * 0.055)
			case sales > 500 && sales <= 1000: fmt.Printf("%.2f", sales * 0.08)
			case sales > 1000 && sales <= 10000: fmt.Printf("%.2f", sales * 0.12)
			case sales > 10000: fmt.Printf("%.2f", sales * 0.145)
		}
	} else {
		fmt.Println("error")
		return
	}
}