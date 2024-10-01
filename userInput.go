package main

import (
	"fmt"
	"strconv"
)

func userInput() int {
	var number string
	var sum int

	fmt.Println("Enter the number or press x to exit:- ")

	for {
		fmt.Scan(&number)

		if number == "x" {
			break
		}

		num, err := strconv.Atoi(number)
		if err != nil {
			fmt.Println("Failed")
			continue
		}
		sum = sum + num

	}
	return sum
}

func main() {
	var result int = userInput()
	fmt.Printf("The sum of the numbers is:- %v", result)
}
