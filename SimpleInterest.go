package main

import (
	"fmt"
)

func simpleInterest(p, r, t int) int {
	return p * r * t / 100
}

func main() {
	var p, r, t int
	fmt.Println("Enter the principal amount")
	fmt.Scan(&p)
	fmt.Println("Enter the rate of interest")
	fmt.Scan(&r)
	fmt.Println("Enter the time period")
	fmt.Scan(&t)
	fmt.Println("The simple interest is", simpleInterest(p, r, t))

}
