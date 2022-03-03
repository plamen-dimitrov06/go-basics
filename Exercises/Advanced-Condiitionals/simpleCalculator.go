package main

import "fmt"

func main() {
	var firstNumber float32
	var secondNumber float32
	var operator string
	fmt.Scanln(&firstNumber)
	fmt.Scanln(&secondNumber)
	fmt.Scanln(&operator)

	if secondNumber == 0 && (operator == "/" || operator == "%") {
		fmt.Printf("Cannot divide %.0f by zero", firstNumber)
		return
	}

	var result float32
	var oddOrEven string
	switch operator {
		case "+":
			result = firstNumber + secondNumber
			if int(result) % 2 == 0 {
				oddOrEven = "even"
			} else {
				oddOrEven = "odd"
			}
			fmt.Printf("%.0f %s %.0f = %.0f - %s", firstNumber, operator, secondNumber, result, oddOrEven)
		case "-":
			result = firstNumber - secondNumber
			if int(result) % 2 == 0 {
				oddOrEven = "even"
			} else {
				oddOrEven = "odd"
			}
			fmt.Printf("%.0f %s %.0f = %.0f - %s", firstNumber, operator, secondNumber, result, oddOrEven)
		case "*":
			result = firstNumber * secondNumber
			if int(result) % 2 == 0 {
				oddOrEven = "even"
			} else {
				oddOrEven = "odd"
			}
			fmt.Printf("%.0f %s %.0f = %.0f - %s", firstNumber, operator, secondNumber, result, oddOrEven)
		case "/":
			result = firstNumber / secondNumber
			fmt.Printf("%.0f %s %.0f = %.2f", firstNumber, operator, secondNumber, result)
		case "%":
			// no the most elegant solution, perhaps the types should be different.
			result = float32(int(firstNumber) % int(secondNumber))
			fmt.Printf("%.0f %s %.0f = %.0f", firstNumber, operator, secondNumber, result)
	}
}