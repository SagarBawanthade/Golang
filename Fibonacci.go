package main

import "fmt"

func fibo(n int) {
	a, b := 0, 1
	var temp int
	for i := 0; i < n; i++ {
		fmt.Print(a, " ")

		temp = a + b
		a = b
		b = temp
	}
}
func main() {
	var n int
	fmt.Println("Enter the number of terms")
	fmt.Scan(&n)
	fibo(n)

}
