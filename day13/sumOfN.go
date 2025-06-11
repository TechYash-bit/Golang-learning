package main

import "fmt"

func main() {
	var method string
	var n int

	// Ask the user which method to use
	fmt.Print("Enter 'loop' or 'formula' to calculate the sum: ")
	fmt.Scan(&method)

	// Ask the user for the upper limit
	fmt.Print("Enter the number up to which you want to calculate the sum: ")
	fmt.Scan(&n)

	sum := 0

	switch method {
	case "loop":
		for i := 1; i <= n; i++ {
			sum += i
		}
		fmt.Println(sum, "→ calculated using loop")
	case "formula":
		sum = n * (n + 1) / 2
		fmt.Println(sum, "→ calculated using formula")
	default:
		fmt.Println("Invalid input. Please enter either 'loop' or 'formula'.")
	}
}
