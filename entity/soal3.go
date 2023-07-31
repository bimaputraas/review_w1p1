package entity

import (
	"fmt"
	"strconv"
)

func iterateNumber(number int, typeNumber string) {
	if typeNumber == "even" {
		fmt.Println("Even number upto your input are:")
		for i := 1; i < number/2+1; i++ {
			fmt.Println(i * 2)
		}
	} else {
		odd := -1
		fmt.Println("Odd number upto your input are:")
		for i := 1; i < (number+3)/2; i++ {
			odd += 2
			fmt.Println(odd)
		}
	}
}

func isPrimeNumber(number int) {
	if (number%2 == 0 || number%3 == 0) && (number != 2 && number != 3) {
		fmt.Println("The number is not a prime number.")
	} else {
		fmt.Println("The number is a prime number.")
	}
}

func Soal3() {
	inputStr := ""
	fmt.Println("Please enter a number:")
	fmt.Scanln(&inputStr)

	inputInt, err := strconv.Atoi(inputStr)
	if err != nil {
		fmt.Println("invalid input")
	}

	if inputInt%2 == 0 {
		fmt.Println("The number is even")
		iterateNumber(inputInt, "even")
	} else {
		fmt.Println("The number is odd")
		iterateNumber(inputInt, "odd")
	}

	isPrimeNumber(inputInt)

}
