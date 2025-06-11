package main

import "fmt"

func main() {
	num := []int{123, 125, 124}
	fmt.Println(num)

	// Allocate space for num1
	num1 := make([]int, len(num)) // Allocate with the same length as num

	// Copying slice from one to another
	fmt.Println("before", num1)
	copy(num1, num)
	fmt.Println("num1", num1)
}
