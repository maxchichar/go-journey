package main

import "fmt"

// add two integers and return their sum
func add(a int, b int) int {
	return a + b
}

func main() {
	sum := add(5, 7)
	fmt.Println("sum:", sum)
}
