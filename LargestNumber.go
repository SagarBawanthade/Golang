package main

import "fmt"

func largestNumber(n1, n2, n3 int) int {
	if n1 > n2 && n1 > n3 {
		return n1
	} else if n2 > n1 && n2 > n3 {
		return n2
	} else {
		return n3
	}
}
func main() {
	var n1, n2, n3 int
	fmt.Println("Enter three numbers")
	fmt.Scan(&n1)
	fmt.Scan(&n2)
	fmt.Scan(&n3)
	fmt.Println("The largest number is", largestNumber(n1, n2, n3))
}
