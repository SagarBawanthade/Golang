//Take 2 numbers as inputs and find their HCF and LCM.

package main

import (
	"fmt"
)

func hcf(a, b int) int {
	if b == 0 {
		return a
	}
	return hcf(b, a%b)
}
func lcm(a, b int) int {
	return (a * b) / hcf(a, b)
}

func main() {
	var a, b int
	fmt.Println("Enter first numbers")
	fmt.Scan(&a)
	fmt.Println("Enter second numbers")
	fmt.Scan(&b)
	fmt.Println("HCF of ", a, " and ", b, " is ", hcf(a, b))
	fmt.Println("LCM of ", a, " and ", b, " is ", lcm(a, b))
}
