package main

import (
	"fmt"
	"slices"
)

func main() {
	original := []int{10, 20, 30, 40, 50}
	fmt.Println("Original slice:     ", original)

	// Clone
	cloned := slices.Clone(original)
	fmt.Println("Cloned slice:       ", cloned)

	// Delete elements from index 1 to 3 (removes 20, 30)
	deleted := slices.Delete(original, 1, 3)
	fmt.Println("After Delete(1,3):  ", deleted)

	// Insert 25, 35 at index 2
	inserted := slices.Insert(original, 2, 25, 35,69)
	fmt.Println("After Insert(2,...):", inserted)

	// Replace elements from index 1 to 3 with 99, 100
	replaced := slices.Replace(original, 1, 3, 99,100)
	fmt.Println("After Replace(1,3): ", replaced)
}
