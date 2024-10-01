package main

import "fmt"

func main() {
	var year int
	fmt.Println("Enter the Year you want to check for Leap Year:-")
	fmt.Scan(&year)

	if year%4 == 0 && year%100 != 0 || year%400 == 0 {
		fmt.Println("The year is a Leap Year")

	} else {
		fmt.Println("The year is not a Leap Year")
	}
}
