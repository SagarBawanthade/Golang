//Take a number as input and print the multiplication table for it.

package main

import "fmt"

func multiplicationNumbers(num int) {
	fmt.Println("Multiplication Table of ", num)
	for i := 1; i <= 10; i++ {
		res := num * i
		fmt.Printf("%v * %v = %v\n", num, i, res)
	}
}

func main() {
	var num int
	fmt.Println("Enter the number you want to multiply: ")
	fmt.Scan(&num)
	multiplicationNumbers(num)
}
