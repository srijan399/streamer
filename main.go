package main

import "fmt"

func main() {
	fmt.Println("Welcome to main")

	var num1 int
	fmt.Println("Enter numbers 1:")
	fmt.Scan(&num1)

	var num2 int
	fmt.Println("Enter numbers 2:")
	fmt.Scan(&num2)

	smaller := findMin(num1, num2)
	fmt.Printf("\nSmaller num is %v\n", smaller)
}

func findMin(n1 int, n2 int) int {
	if n1 < n2 {
		return n1
	} else {
		return n2
	}
}
