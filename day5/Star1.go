package main

import "fmt"

func main() {
	fmt.Println("Enter the number of rows:")
	var row int
	fmt.Scan(&row)

	for i := 0; i < row; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
