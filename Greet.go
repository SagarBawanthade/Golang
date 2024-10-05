package main

import "fmt"

func greet(name string) string {
	return "Hello, " + name + "!"
}

func main() {
	var name string
	fmt.Println("Enter your name")
	fmt.Scan(&name)
	fmt.Println(greet(name))
}
