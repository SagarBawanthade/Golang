//Take two numbers and print the sum of both.

package main

import "fmt"

func sumOfTwoNumbers(num1, num2, sum int) int {
	sum = num1 + num2
	return sum
}

func main() {
	var num1, num2, sum, result int
	fmt.Println("Enter the first number:-")
	fmt.Scan(&num1)
	fmt.Println("Enter the second number:-")
	fmt.Scan(&num2)
	result = sumOfTwoNumbers(num1, num2, sum)
	fmt.Println("The sum of the two numbers is:-", result)
}
