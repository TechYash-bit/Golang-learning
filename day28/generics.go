package main

import "fmt"

func swap[T any](a, b T) (T, T) {
	return b, a
}
func main() {
	a, b := "yash", "harne"

	a, b = swap(a, b)

	fmt.Println(a," ",b)

}