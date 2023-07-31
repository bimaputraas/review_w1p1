package entity

import (
	"fmt"
	"os"
	"strconv"
)

func calculator(operator string, number1, number2 float64) {
	output := float64(0)
	if operator == "add" {
		output = number1 + number2
	} else if operator == "sub" {
		output = number1 - number2
	} else if operator == "mul" {
		output = number1 * number2
	} else if operator == "div" {
		output = number1 / number2
	}
	fmt.Printf("Result: %.2f\n", output)
}

func Soal2() {
	if len(os.Args) != 4 {
		fmt.Println("Invalid Input")
	} else {
		operator := os.Args[1]
		number1Str := os.Args[2]
		number2Str := os.Args[3]

		number1, _ := strconv.Atoi(number1Str)
		number2, _ := strconv.Atoi(number2Str)
		calculator(operator, float64(number1), float64(number2))
	}

}
